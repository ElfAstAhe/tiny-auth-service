package postgres

import (
	"context"
	"database/sql"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

const (
	sqlUserRolesAdminFind string = `
select
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
where
    ur.user_id = $1
and ur.role_id = $2
`
	sqlUserRolesAdminListAll string = `
select
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
where
    ur.user_id = $1
`
	sqlUserRolesAdminListAllByOwners string = `
select
    ur.user_id,
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
where
    ur.user_id = any($1)
`
	sqlUserRolesAdminCreate string = `
insert into user_roles (
    user_id,
    role_id
)
values ($1, $2)
`
	sqlUserRolesAdminDelete string = `
delete from user_roles where user_id = $1
`
)

type UserRolesAdminPgRepository struct {
	*repository.BaseOwnedRepository[*domain.Role, string, string]
}

func NewUserRolesAdminPgRepository(exec db.Executor, errDecipher db.ErrorDecipher) (*UserRolesAdminPgRepository, error) {
	res := &UserRolesAdminPgRepository{}
	// query builders
	queryBuilders := repository.NewBaseOwnedQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlUserRolesAdminFind
		}).
		WithListAll(func() string {
			return sqlUserRolesAdminListAll
		}).
		WithListAllByOwners(func() string {
			return sqlUserRolesAdminListAllByOwners
		}).
		WithCreate(func() string {
			return sqlUserRolesAdminCreate
		}).
		WithDelete(func() string {
			return sqlUserRolesAdminDelete
		}).
		Build()

	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyRole).
		WithValidateCreate(res.validateCreate).
		WithCreator(res.creator).
		Build()
	if err != nil {
		return nil, err
	}

	base, err := repository.NewBaseOwnedRepository[*domain.Role, string, string](
		exec,
		errDecipher,
		repository.NewEntityInfo("user_roles", "UserRole"),
		queryBuilders,
		callbacks,
		repository.LinkStrategyManyToMany,
		nil,
	)

	res.BaseOwnedRepository = base

	return res, nil
}

func (ura *UserRolesAdminPgRepository) entityScanner(scanner repository.Scannable, dest *domain.Role, params ...any) error {
	if len(params) == 0 {
		return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	}

	return scanner.Scan(&params[0], &dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (ura *UserRolesAdminPgRepository) validateCreate(role *domain.Role, params ...any) error {
	if role == nil {
		return errs.NewInvalidArgumentError("role", "role is nil")
	}
	if strings.TrimSpace(role.ID) == "" {
		return errs.NewInvalidArgumentError("role", "role id is empty")
	}
	if len(params) == 0 {
		return errs.NewInvalidArgumentError("params", "params is empty")
	}
	userID, ok := params[0].(string)
	if !ok {
		return errs.NewInvalidArgumentError("userID", "must be a string")
	}
	if strings.TrimSpace(userID) == "" {
		return errs.NewInvalidArgumentError("userID", "must not be empty")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, ura.GetQueryBuilders().GetCreate()(), params[0], entity.ID), nil
}

func (ura *UserRolesAdminPgRepository) ValidateDeleteAll(ownerID string) error {
	if strings.TrimSpace(ownerID) == "" {
		return errs.NewInvalidArgumentError("ownerID", "must not be empty")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) ValidateDelete(ownerID string) error {
	if strings.TrimSpace(ownerID) == "" {
		return errs.NewInvalidArgumentError("ownerID", "must not be empty")
	}

	return nil
}
