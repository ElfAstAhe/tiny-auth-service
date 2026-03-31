package mapper

import (
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

func MapAuthToDTO(token string, refreshToken string) *dto.LoggedInDTO {
	return &dto.LoggedInDTO{
		Token:        token,
		RefreshToken: refreshToken,
	}
}
