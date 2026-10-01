package grpc

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MapProfileDTOToGRPC transforms a facade ProfileDTO layer structure into a strongly-typed wire-ready gRPC pb.ProfileResponse package payload.
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

// MapChangePasswordGRPCToDTO converts an inbound gRPC pb.ChangePasswordRequest network criteria payload into a decoupled internal facade.ChangePasswordDTO.
func MapChangePasswordGRPCToDTO(req *pb.ChangePasswordRequest) *dto.ChangePasswordDTO {
	if req == nil {
		return nil
	}

	return &dto.ChangePasswordDTO{
		OldPassword: req.GetOldPassword(),
		NewPassword: req.GetNewPassword(),
	}
}

// MapChangedKeysDTOToGRPC translates internal facade ChangedKeysDTO records back into public cryptographically signed gRPC pb.ChangeKeysResponse payloads.
func MapChangedKeysDTOToGRPC(resp *dto.ChangedKeysDTO) *pb.ChangeKeysResponse {
	if resp == nil {
		return nil
	}

	return pb.ChangeKeysResponse_builder{
		PublicKey: &resp.PublicKey,
	}.Build()
}
