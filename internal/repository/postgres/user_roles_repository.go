package postgres

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

const (
	sqlUserRolesList string = `
select
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join roles r
`
)

type UserRolesPgRepository struct {
	exec        db.Executor
	errDecipher db.ErrorDecipher
	info        *repository.EntityInfo
}

func NewUserRolesPgRepository(executor db.Executor, errDecipher db.ErrorDecipher) *UserRolesPgRepository {
	return &UserRolesPgRepository{
		exec:        executor,
		errDecipher: errDecipher,
		info:        repository.NewEntityInfo("user_roles", "UserRole"),
	}
}

func (urr *UserRolesPgRepository) ListRoles(ctx context.Context, userID string) ([]*domain.Role, error) {

}
