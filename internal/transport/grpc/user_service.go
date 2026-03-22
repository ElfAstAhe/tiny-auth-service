package grpc

import (
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

type UserGRPCService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserGRPCService() *UserGRPCService {
	return &UserGRPCService{}
}
