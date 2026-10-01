package grpc

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MapRoleDTOToGRPC transforms a facade RoleDTO layer structure into a strongly-typed wire-ready gRPC pb.Role message payload.
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

// MapRoleGRPCToDTO converts an inbound administrative gRPC pb.Role network criteria payload into a decoupled internal facade.RoleDTO.
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

// MapRoleDTOsToGRPC converts a slice array of facade RoleDTO pointers into a decoupled slice collection array of gRPC pb.Role entities.
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

// MapRoleGRPCsToDTO reverses a collection array of gRPC pb.Role payloads back into internal facade RoleDTO component structures.
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
