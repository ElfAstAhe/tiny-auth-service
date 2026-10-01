package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// MapUserModelToDTO transforms a domain User aggregate entity directly into a flattened, data-transfer-safe dto.UserDTO blueprint structure.
func MapUserModelToDTO(model *domain.User) *dto.UserDTO {
	if model == nil {
		return nil
	}

	return &dto.UserDTO{
		ID:           model.ID,
		Name:         model.Name,
		Type:         model.Type,
		PasswordHash: model.PasswordHash,
		PublicKey:    model.PublicKey,
		PrivateKey:   model.PrivateKey,
		Active:       model.Active,
		Deleted:      model.Deleted,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		Roles:        MapRolesModelToDTO(model.Roles),
	}
}

// MapUserModelsToDTO converts a slice array of domain User pointers into a decoupled slice collection array of dto.UserDTO entities.
func MapUserModelsToDTO(models []*domain.User) []*dto.UserDTO {
	if len(models) == 0 {
		return make([]*dto.UserDTO, 0)
	}

	res := make([]*dto.UserDTO, 0, len(models))
	for _, model := range models {
		res = append(res, MapUserModelToDTO(model))
	}

	return res
}

// MapUserDTOToModel reverses a data-transfer object payload schema, re-assembling a structured domain.User core execution aggregate entity.
func MapUserDTOToModel(user *dto.UserDTO) *domain.User {
	if user == nil {
		return nil
	}

	return &domain.User{
		ID:           user.ID,
		Name:         user.Name,
		Type:         user.Type,
		PasswordHash: user.PasswordHash,
		PublicKey:    user.PublicKey,
		PrivateKey:   user.PrivateKey,
		Active:       user.Active,
		Deleted:      user.Deleted,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Roles:        MapRolesDTOToModel(user.Roles),
	}
}
