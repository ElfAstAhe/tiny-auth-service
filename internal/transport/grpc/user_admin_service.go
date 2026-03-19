package grpc

import (
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

type UserAdminGRPCService struct {
	pb.AdminUsersServiceServer
}

func NewUserAdminGRPCService() *UserAdminGRPCService {
	return &UserAdminGRPCService{}
}
