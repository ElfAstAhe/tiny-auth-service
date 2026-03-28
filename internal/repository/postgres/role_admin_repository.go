package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

const (
	sqlRoleAdminFind string = `
select
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
from
    roles
where
    id = $1
`
	sqlRoleAdminFindByName string = `
select
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
from
    roles
where
    name = $1
`
	sqlRoleAdminList string = `
select
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
from
    roles
order by
    id asc
offset $2
limit $1
`
	sqlRoleAdminCreate string = `
insert into roles (
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
)
values ($1, $2, $3, $4, $5, $6)
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleAdminChange string = `
update
    roles
set
    name = $2,
    description = $3,
    deleted = $4,
    updated_at = $5
where
    id = $1
`
	sqlRoleAdminDelete string = `
delete
from
    roles
where
    id = $1
`
)

type RoleAdminPgRepository struct {
	*repository.BaseCRUDRepository[*domain.Role, string]
	userRolesRepo domain.UserRolesAdminRepository
}

func NewRoleAdminPgRepository(executor db.Executor, decipher db.ErrorDecipher) (*RoleAdminPgRepository, error) {
	// new instance
	res := &RoleAdminPgRepository{}
	// sql builders
	queryBuilders := repository.NewBaseCRUDQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlRoleAdminFind
		}).
		WithList(func() string {
			return sqlRoleAdminList
		}).
		WithCreate(func() string {
			return sqlRoleAdminCreate
		}).
		WithChange(func() string {
			return sqlRoleAdminChange
		}).
		WithDelete(func() string {
			return sqlRoleAdminDelete
		}).
		Build()
	// callbacks
	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyRole).
		WithValidateCreate(res.validateCreate).
		WithBeforeCreate(res.beforeCreate).
		WithCreator(res.creator).
		WithValidateChange(res.validateChange).
		WithBeforeChange(res.beforeChange).
		WithChanger(res.changer).
		Build()
	// base CRUD
	base, err := repository.NewBaseCRUDRepository[*domain.Role, string](
		executor,
		decipher,
		repository.NewEntityInfo("roles", "Role"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create RolePgRepository", err)
	}

	res.BaseCRUDRepository = base

	return res, nil
}

func (rar *RoleAdminPgRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	if name == "" {
		return nil, errs.NewInvalidArgumentError("name", "cannot be empty")
	}

	return rar.GetHelper().Get(ctx, sqlRoleAdminFindByName, name)
}

func (rar *RoleAdminPgRepository) entityScanner(scanner repository.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (rar *RoleAdminPgRepository) validateCreate(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateCreate()
}

func (rar *RoleAdminPgRepository) beforeCreate(entity *domain.Role, params ...any) error {
	if err := entity.ValidateCreate(); err != nil {
		return errs.NewDalError("RoleAdminPgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

func (rar *RoleAdminPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rar.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.Description, entity.CreatedAt, entity.UpdatedAt), nil
}

func (rar *RoleAdminPgRepository) validateChange(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateChange()
}

func (rar *RoleAdminPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rar.GetQueryBuilders().GetChange()(), entity.ID, entity.Name, entity.Description, entity.UpdatedAt), nil
}

func (rar *RoleAdminPgRepository) beforeChange(entity *domain.Role, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("RoleAdminPgRepository.beforeChange", "before change entity", err)
	}

	return nil
}
