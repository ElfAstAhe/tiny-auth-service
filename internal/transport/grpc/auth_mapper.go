package grpc

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

func MapLoginRespDTOToGRPC(resp *dto.LoggedInDTO) *pb.AuthLoginResponse {
	if resp == nil {
		return nil
	}

	return pb.AuthLoginResponse_builder{
		Token:        &resp.Token,
		RefreshToken: &resp.RefreshToken,
	}.Build()
}

func MapLoginReqGRPCToDTO(req *pb.AuthLoginRequest) *dto.LoginDTO {
	if req == nil {
		return nil
	}

	return &dto.LoginDTO{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
}
