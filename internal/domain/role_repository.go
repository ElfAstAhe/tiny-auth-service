package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type RoleRepository interface {
	domain.CRUDRepository[*Role, string]

	FindByName(context.Context, string) (*Role, error)
}
