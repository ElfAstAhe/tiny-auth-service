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

func NewUserGRPCService(userFacade facade.UserFacade) *UserGRPCService {
	return &UserGRPCService{
		userFacade: userFacade,
	}
}

// Profile get user profile info
func (us *UserGRPCService) Profile(context.Context, *emptypb.Empty) (*pb.ProfileResponse, error) {

}

// ChangePassword changes user password
func (us *UserGRPCService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*emptypb.Empty, error) {

}

// ChangeKeys generate new RSA key pair
func (us *UserGRPCService) ChangeKeys(context.Context, *emptypb.Empty) (*pb.ChangeKeysResponse, error) {

}
