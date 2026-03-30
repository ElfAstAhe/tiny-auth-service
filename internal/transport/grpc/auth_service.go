package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

type AuthGRPCService struct {
	pb.UnimplementedAuthServiceServer
	authFacade facade.AuthFacade
}

var _ pb.AuthServiceServer = (*AuthGRPCService)(nil)

func NewAuthGRPCService(authFacade facade.AuthFacade) *AuthGRPCService {
	return &AuthGRPCService{
		authFacade: authFacade,
	}
}

func (as *AuthGRPCService) Login(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	dtoRes, err := as.authFacade.Login(ctx, MapLoginReqGRPCToDTO(req))
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return MapLoginRespDTOToGRPC(dtoRes), nil
}
