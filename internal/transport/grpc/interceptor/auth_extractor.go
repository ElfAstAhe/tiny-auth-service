package interceptor

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// wrappedStream structures an internal proxy interceptor mapping table to override context parameters
// inside live long-lived gRPC streaming network data connections safely.
type wrappedStream struct {
	grpc.ServerStream                 // Embedded standard core library server stream structure interface handle
	ctx               context.Context // Target identity enriched context payload embedded within current connection scope
}

// Context returns the newly assigned structural context carrying verified authorization token claims.
func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

// AuthExtractor manages authentication verification checkpoints across unary and streaming gRPC network endpoints.
// It tracks public route access allowances and injects security token metadata into runtime contexts.
type AuthExtractor struct {
	authHelper auth.Helper         // Framework security utility responsible for decoding incoming transport context metadata
	log        logger.Logger       // Structured logging handle isolating security interceptor event records
	nonSecure  map[string]struct{} // Read-only evaluation map pinning registration identifiers for unauthenticated methods
}

// NewAuthExtractor acts as a factory constructor setting up target non-secure route maps and logging prefixes.
func NewAuthExtractor(authHelper auth.Helper, logger logger.Logger) *AuthExtractor {
	return &AuthExtractor{
		authHelper: authHelper,
		log:        logger.GetLogger("gRPC-Auth-Extractor"),
		nonSecure: map[string]struct{}{
			pb.AuthService_Login_FullMethodName:       {},
			pb.AuthService_LoginSimple_FullMethodName: {},
		},
	}
}

// UnaryServerInterceptor captures single request-response sequences to validate signatures and mount valid claims payloads.
func (ae *AuthExtractor) UnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	ae.log.Debugf("UnaryServerInterceptor start with req: [%v]", req)
	defer ae.log.Debug("UnaryServerInterceptor finish")

	// ignorance
	if ae.isNonSecure(info.FullMethod) {
		return handler(ctx, req)
	}

	subj, err := ae.authHelper.SubjectFromGRPCContext(ctx)
	if err != nil {
		ae.log.Errorf("AuthExtractor.UnaryServerInterceptor failed with error: [%v]", err)

		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	secureCtx := auth.WithSubject(ctx, subj)

	return handler(secureCtx, req)
}

// StreamServerInterceptor handles continuous data pipeline connections, wrapping interfaces into context adapters dynamically.
func (ae *AuthExtractor) StreamServerInterceptor(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ae.log.Debugf("StreamServerInterceptor start: method [%v]", info.FullMethod)
	defer ae.log.Debugf("StreamServerInterceptor finish: method [%v]", info.FullMethod)

	// ignorance
	if ae.isNonSecure(info.FullMethod) {
		return handler(srv, stream)
	}

	subj, err := ae.authHelper.SubjectFromGRPCContext(stream.Context())
	if err != nil {
		ae.log.Errorf("AuthExtractor.StreamServerInterceptor failed with error: [%v]", err)

		return status.Error(codes.Unauthenticated, err.Error())
	}

	wrapped := &wrappedStream{
		ServerStream: stream,
		ctx:          auth.WithSubject(stream.Context(), subj),
	}

	return handler(srv, wrapped)
}

// isNonSecure checks if incoming structural lookup endpoint keys bypass standard security verification rules.
func (ae *AuthExtractor) isNonSecure(method string) bool {
	_, ok := ae.nonSecure[method]

	return ok
}
