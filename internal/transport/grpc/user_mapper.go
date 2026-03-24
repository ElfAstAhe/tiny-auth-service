package grpc

import (
    "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
    pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
)

func MapProfileDTOToGRPC(profile *dto.ProfileDTO) *pb.ProfileResponse {
    if profile == nil {
        return nil
    }

    return &pb.ProfileResponse_builder{
        Info:
    }
}