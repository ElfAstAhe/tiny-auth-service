package grpc

import (
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

type AuthGRPCService struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthGRPCService() *AuthGRPCService {
	return &AuthGRPCService{}
}
