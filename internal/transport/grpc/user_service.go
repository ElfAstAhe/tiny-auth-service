package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserGRPCService struct {
	pb.UnimplementedUserServiceServer
	userFacade facade.UserFacade
}

var _ pb.UserServiceServer = (*UserGRPCService)(nil)

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
