package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// MapRoleModelToDTO transforms a domain Role aggregate entity directly into a flattened, data-transfer-safe dto.RoleDTO blueprint structure.
func MapRoleModelToDTO(model *domain.Role) *dto.RoleDTO {
	if model == nil {
		return nil
	}

	return &dto.RoleDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Deleted:     model.Deleted,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// MapRolesModelToDTO converts a slice array of domain Role pointers into a decoupled slice collection array of dto.RoleDTO entities.
func MapRolesModelToDTO(models []*domain.Role) []*dto.RoleDTO {
	if len(models) == 0 {
		return make([]*dto.RoleDTO, 0)
	}

	res := make([]*dto.RoleDTO, 0, len(models))
	for _, model := range models {
		res = append(res, MapRoleModelToDTO(model))
	}

	return res
}

// MapRoleDTOToModel reverses a data-transfer object payload schema, re-assembling a structured domain.Role core execution aggregate entity.
func MapRoleDTOToModel(role *dto.RoleDTO) *domain.Role {
	if role == nil {
		return nil
	}

	return &domain.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Deleted:     role.Deleted,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

// MapRolesDTOToModel reverses a collection array of data-transfer objects payloads back into a slice of structured domain.Role core execution aggregate entities.
func MapRolesDTOToModel(roles []*dto.RoleDTO) []*domain.Role {
	if len(roles) == 0 {
		return make([]*domain.Role, 0)
	}

	res := make([]*domain.Role, 0, len(roles))
	for _, role := range roles {
		res = append(res, MapRoleDTOToModel(role))
	}

	return res
}
