package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type UserRepository interface {
	domain.CRUDRepository[*User, string]

	FindByName(ctx context.Context, login string) (*User, error)
}
