package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

func MapUserModelToDTO(model *domain.User) *dto.UserDTO {
	if model == nil {
		return nil
	}

	return &dto.UserDTO{
		ID:           model.ID,
		Name:         model.Name,
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

func MapUserDTOToModel(user *dto.UserDTO) *domain.User {
	if user == nil {
		return nil
	}

	return &domain.User{
		ID:           user.ID,
		Name:         user.Name,
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
