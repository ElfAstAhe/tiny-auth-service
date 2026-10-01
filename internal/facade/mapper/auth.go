package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// MapAuthToDTO encapsulates transformation mechanics converting raw signed token string sequences into a structured data transfer dto.LoggedInDTO payload object.
func MapAuthToDTO(token string, refreshToken string) *dto.LoggedInDTO {
	return &dto.LoggedInDTO{
		Token:        token,
		RefreshToken: refreshToken,
	}
}
