package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// MapUserModelToProfileDTO transforms a domain User aggregate entity directly into a decoupled, data-transfer-safe dto.ProfileDTO blueprint structure.
func MapUserModelToProfileDTO(model *domain.User) *dto.ProfileDTO {
	if model == nil {
		return nil
	}

	res := &dto.ProfileDTO{
		ID:        model.ID,
		Name:      model.Name,
		Type:      model.Type,
		PublicKey: model.PublicKey,
		Active:    model.Active,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Roles:     MapRolesModelToNames(model.Roles),
	}

	return res
}

// MapRolesModelToNames extracts and transforms a slice array of domain Role pointers into a flat slice collection array of string names.
func MapRolesModelToNames(models []*domain.Role) []string {
	res := make([]string, 0, len(models))
	for _, model := range models {
		res = append(res, model.Name)
	}

	return res
}

// MapKeysToDTO encapsulates transformation routines wrapping a public key string literal into a structured dto.ChangedKeysDTO payload wrapper.
func MapKeysToDTO(publicKey string) *dto.ChangedKeysDTO {
	return &dto.ChangedKeysDTO{
		PublicKey: publicKey,
	}
}
