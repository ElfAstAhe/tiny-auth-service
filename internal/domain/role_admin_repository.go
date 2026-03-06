package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type RoleAdminRepository interface {
	domain.Repository[*Role, string]

	FindByName(ctx context.Context, login string) (*Role, error)
}
