package grpc

import (
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

type RoleAdminGRPCService struct {
	pb.AdminRolesServiceServer
}

func NewRoleAdminGRPCService() *RoleAdminGRPCService {
	return &RoleAdminGRPCService{}
}
