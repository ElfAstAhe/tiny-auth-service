package facade

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type AuthFacade interface {
	Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error)
}
