package grpc

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapRoleDTOToGRPC(instance *dto.RoleDTO) *pb.Role {
	if instance == nil {
		return nil
	}

	return pb.Role_builder{
		Id:          &instance.ID,
		Name:        &instance.Name,
		Description: &instance.Description,
		Deleted:     &instance.Deleted,
		CreatedAt:   timestamppb.New(instance.CreatedAt),
		UpdatedAt:   timestamppb.New(instance.UpdatedAt),
	}.Build()
}

func MapRoleGRPCToDTO(instance *pb.Role) *dto.RoleDTO {
	if instance == nil {
		return nil
	}

	return &dto.RoleDTO{
		ID:          instance.GetId(),
		Name:        instance.GetName(),
		Description: instance.GetDescription(),
		Deleted:     instance.GetDeleted(),
		CreatedAt:   instance.GetCreatedAt().AsTime(),
		UpdatedAt:   instance.GetUpdatedAt().AsTime(),
	}
}

func MapRoleDTOsToGRPC(roles []*dto.RoleDTO) []*pb.Role {
	if len(roles) == 0 {
		return make([]*pb.Role, 0)
	}

	res := make([]*pb.Role, 0, len(roles))
	for _, role := range roles {
		res = append(res, MapRoleDTOToGRPC(role))
	}

	return res
}

func MapRoleGRPCsToDTO(roles []*pb.Role) []*dto.RoleDTO {
	if len(roles) == 0 {
		return make([]*dto.RoleDTO, 0)
	}

	res := make([]*dto.RoleDTO, 0, len(roles))
	for _, role := range roles {
		res = append(res, MapRoleGRPCToDTO(role))
	}

	return res
}
