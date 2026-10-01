package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserGRPCService implements the pb.UserServiceServer gRPC interface network delivery boundary.
// It translates incoming high-performance binary Protobuf streams into structured Facade DTO objects.
type UserGRPCService struct {
	pb.UnimplementedUserServiceServer                   // Embedded fallback forward-compatibility enforcement handle
	userFacade                        facade.UserFacade // Application boundary facade router executing coordinated business scenarios
}

// Compile-time interface compliance verification
var _ pb.UserServiceServer = (*UserGRPCService)(nil)

// NewUserGRPCService acts as a factory constructor mounting required facade layer dependencies.
func NewUserGRPCService(userFacade facade.UserFacade) *UserGRPCService {
	return &UserGRPCService{
		userFacade: userFacade,
	}
}

// Profile get user profile info
func (us *UserGRPCService) Profile(ctx context.Context, req *emptypb.Empty) (*pb.ProfileResponse, error) {
	res, err := us.userFacade.Profile(ctx)
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return MapProfileDTOToGRPC(res), nil
}

// ChangePassword changes user password
func (us *UserGRPCService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*emptypb.Empty, error) {
	// PROTO PAYLOAD NOTICE: The inbound request 'req' is passed directly into the transformer mapper.
	// Ensure that 'MapChangePasswordGRPCToDTO' executes a strict nil-pointer check to defend the runtime from potential panic states under malicious grpc payloads.
	err := us.userFacade.ChangePassword(ctx, MapChangePasswordGRPCToDTO(req))
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}

// ChangeKeys generate new RSA key pair
func (us *UserGRPCService) ChangeKeys(ctx context.Context, req *emptypb.Empty) (*pb.ChangeKeysResponse, error) {
	res, err := us.userFacade.ChangeKeys(ctx)
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return MapChangedKeysDTOToGRPC(res), nil
}
