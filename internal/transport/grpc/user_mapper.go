package grpc

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapProfileDTOToGRPC(profile *dto.ProfileDTO) *pb.ProfileResponse {
	if profile == nil {
		return nil
	}

	return pb.ProfileResponse_builder{
		Id:        &profile.ID,
		Name:      &profile.Name,
		UserType:  &profile.Type,
		PublicKey: &profile.PublicKey,
		Active:    &profile.Active,
		CreatedAt: timestamppb.New(profile.CreatedAt),
		UpdatedAt: timestamppb.New(profile.UpdatedAt),
		Roles:     profile.Roles,
	}.Build()
}

func MapChangePasswordGRPCToDTO(req *pb.ChangePasswordRequest) *dto.ChangePasswordDTO {
	if req == nil {
		return nil
	}

	return &dto.ChangePasswordDTO{
		OldPassword: req.GetOldPassword(),
		NewPassword: req.GetNewPassword(),
	}
}

func MapChangedKeysDTOToGRPC(resp *dto.ChangedKeysDTO) *pb.ChangeKeysResponse {
	if resp == nil {
		return nil
	}

	return pb.ChangeKeysResponse_builder{
		PublicKey: &resp.PublicKey,
	}.Build()
}
