package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

func MapUserModelToProfileDTO(model *domain.User) *dto.ProfileDTO {
	if model == nil {
		return nil
	}

	res := &dto.ProfileDTO{
		ID:        model.ID,
		Name:      model.Name,
		PublicKey: model.PublicKey,
		Active:    model.Active,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Roles:     MapRolesModelToNames(model.Roles),
	}

	return res
}

func MapRolesModelToNames(models []*domain.Role) []string {
	res := make([]string, 0, len(models))
	for _, model := range models {
		res = append(res, model.Name)
	}

	return res
}

func MapKeysToDTO(publicKey string) *dto.ChangedKeysDTO {
	return &dto.ChangedKeysDTO{
		PublicKey: publicKey,
	}
}
