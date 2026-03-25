package grpc

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapUserDTOToGRPC(instance *dto.UserDTO) *pb.User {
	if instance == nil {
		return nil
	}

	return pb.User_builder{
		Id:           &instance.ID,
		Name:         &instance.Name,
		PasswordHash: &instance.PasswordHash,
		PublicKey:    &instance.PublicKey,
		PrivateKey:   &instance.PrivateKey,
		Active:       &instance.Active,
		Deleted:      &instance.Deleted,
		CreatedAt:    timestamppb.New(instance.CreatedAt),
		UpdatedAt:    timestamppb.New(instance.UpdatedAt),
		Roles:        MapRoleDTOsToGRPC(instance.Roles),
	}.Build()
}

func MapUserGRPCToDTO(instance *pb.User) *dto.UserDTO {
	if instance == nil {
		return nil
	}

	return &dto.UserDTO{
		ID:           instance.GetId(),
		Name:         instance.GetName(),
		PasswordHash: instance.GetPasswordHash(),
		PublicKey:    instance.GetPublicKey(),
		PrivateKey:   instance.GetPrivateKey(),
		Active:       instance.GetActive(),
		Deleted:      instance.GetDeleted(),
		Roles:        MapRoleGRPCsToDTO(instance.GetRoles()),
	}
}

func MapUserDTOsToGRPC(instances []*dto.UserDTO) []*pb.User {
	if len(instances) == 0 {
		return make([]*pb.User, 0)
	}

	res := make([]*pb.User, 0, len(instances))
	for _, instance := range instances {
		res = append(res, MapUserDTOToGRPC(instance))
	}

	return res
}
