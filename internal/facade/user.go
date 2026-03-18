package facade

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type UserFacade interface {
	Profile(ctx context.Context) (*dto.UserDTO, error)
	ChangePassword(ctx context.Context, oldPassword, newPassword string) error
	ChangeKeys(ctx context.Context) (*dto.ChangedKeysDTO, error)
}
