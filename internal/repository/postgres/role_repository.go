package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// SQL запросы
const (
	sqlRoleFind string = `
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
	sqlRoleFindByName string = `
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
	sqlRoleList string = `
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
	sqlRoleCreate string = `
insert into roles (
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
)
values($1, $2, $3, false, $4, $5)
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleChange string = `
update
    roles
set
    name = $2,
    description = $3,
    updated_at = $4
where
    id = $1
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleDelete string = `
update
    roles
set
    deleted = true
where
    id = $1
`
)

type RolePgRepository struct {
	*repository.BaseRepository[*domain.Role, string]
}

func NewRolePgRepository(executor db.Executor, decipher db.ErrorDecipher) (*RolePgRepository, error) {
	// new instance
	res := &RolePgRepository{}
	// sql builders
	queryBuilders := repository.NewBaseQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlRoleFind
		}).
		WithList(func() string {
			return sqlRoleList
		}).
		WithCreate(func() string {
			return sqlRoleCreate
		}).
		WithChange(func() string {
			return sqlRoleChange
		}).
		WithDelete(func() string {
			return sqlRoleDelete
		}).
		Build()
	// callbacks
	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyRole).
		WithAfterFind(res.afterFind).
		WithAfterListYield(res.afterListYield).
		WithValidateCreate(res.validateCreate).
		WithBeforeCreate(res.beforeCreate).
		WithCreator(res.creator).
		WithValidateChange(res.validateChange).
		WithBeforeChange(res.beforeChange).
		WithChanger(res.changer).
		Build()
	// base CRUD
	base, err := repository.NewBaseRepository[*domain.Role, string](
		executor,
		decipher,
		repository.NewEntityInfo("roles", "Role"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create RolePgRepository", err)
	}

	res.BaseRepository = base

	return res, nil
}

func (rr *RolePgRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	querier := rr.GetExecutor().GetQuerier(ctx)

	row := querier.QueryRowContext(ctx, sqlRoleFindByName, name)

	res := rr.GetCallbacks().NewEntityFactory()

	err := rr.GetCallbacks().EntityScanner(row, res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewDalNotFoundError(rr.GetInfo().Entity, name, err)
		}

		return nil, errs.NewDalError("RolePgRepository.FindByName", "get row", err)
	}

	if rr.GetCallbacks().AfterFind != nil {
		return rr.GetCallbacks().AfterFind(res)
	}

	return res, nil
}

func (rr *RolePgRepository) entityScanner(scanner repository.Scannable, dest *domain.Role) error {
	return scanner.Scan(&dest.ID, dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (rr *RolePgRepository) afterFind(entity *domain.Role) (*domain.Role, error) {
	if entity.IsDeleted() {
		return nil, errs.NewDalSoftDeletedError(rr.GetInfo().Entity, entity.GetID())
	}

	return entity, nil
}

func (rr *RolePgRepository) afterListYield(entity *domain.Role) (*domain.Role, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(rr.GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

func (rr *RolePgRepository) validateCreate(entity *domain.Role) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateCreate()
}

func (rr *RolePgRepository) beforeCreate(entity *domain.Role) error {
	if err := entity.ValidateCreate(); err != nil {
		return errs.NewDalError("RolePgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

func (rr *RolePgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rr.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.Description, entity.CreatedAt, entity.UpdatedAt), nil
}

func (rr *RolePgRepository) validateChange(entity *domain.Role) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateChange()
}

func (rr *RolePgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.Role) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rr.GetQueryBuilders().GetChange()(), entity.ID, entity.Name, entity.Description, entity.UpdatedAt), nil
}

func (rr *RolePgRepository) beforeChange(entity *domain.Role) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("RolePgRepository.beforeChange", "before change entity", err)
	}

	return nil
}

func (rr *RolePgRepository) Close() error {
	return nil
}
