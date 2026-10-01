package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

// AuthGRPCService implements the pb.AuthServiceServer gRPC interface network delivery boundary.
// It translates incoming high-performance binary Protobuf authentication streams into structured Facade DTO criteria targets.
type AuthGRPCService struct {
	pb.UnimplementedAuthServiceServer                   // Embedded fallback forward-compatibility enforcement handle
	authFacade                        facade.AuthFacade // Application boundary facade router executing coordinated credentials evaluation scenarios
}

// Compile-time interface compliance verification
var _ pb.AuthServiceServer = (*AuthGRPCService)(nil)

// NewAuthGRPCService acts as a factory constructor mounting required authentication facade layer dependencies.
func NewAuthGRPCService(authFacade facade.AuthFacade) *AuthGRPCService {
	return &AuthGRPCService{
		authFacade: authFacade,
	}
}

// Login executes a full security verification scenario matching credentials parsed from the gRPC wire data.
func (as *AuthGRPCService) Login(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	dtoRes, err := as.authFacade.Login(ctx, MapLoginReqGRPCToDTO(req))
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return MapLoginRespDTOToGRPC(dtoRes), nil
}

// LoginSimple executes a simplified credentials evaluation sequence bypassing heavy cryptographic private RSA key operations.
func (as *AuthGRPCService) LoginSimple(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	dtoRes, err := as.authFacade.LoginSimple(ctx, MapLoginReqGRPCToDTO(req))
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return MapLoginRespDTOToGRPC(dtoRes), nil
}
