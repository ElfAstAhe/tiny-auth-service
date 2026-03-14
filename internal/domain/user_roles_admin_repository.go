package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type UserRolesAdminRepository interface {
	domain.OwnedRepository[*Role, string, string]

	DeleteAllOwned(ctx context.Context, ownedID string) error
}
