package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type UserAdminRepository interface {
	domain.CRUDRepository[*User, string]

	FindByName(context.Context, string) (*User, error)
}
