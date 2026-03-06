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

const (
	sqlUserAdminFind string = `
select
    id,
    name,
    password_hash,
    active,
    deleted,
    created_at,
    updated_at
from
    users
where
    id = $1
`
	sqlUserAdminFindByName string = `
select
    id,
    name,
    password_hash,
    active,
    deleted,
    created_at,
    updated_at
from
    users
where
    name = $1
`
	sqlUserAdminList string = `
select
    id,
    name,
    password_hash,
    active,
    deleted,
    created_at,
    updated_at
from
    users
order by
    id asc
offset $2
limit $1
`
	sqlUserAdminCreate string = `
insert into users (
    id,
    name,
    password_hash,
    active,
    deleted,
    created_at,
    updated_at
)
values ($1, $2, $3, $4, $5, $6, $7)
returning id, name, password_hash, active, deleted, created_at, updated_at
`
	sqlUserAdminChange string = `
update
    users
set
    name = $2,
    password_hash = $3,
    active = $4,
    deleted = $5,
    updated_at = $6
where
    id = $1
returning id, name, password_hash, active, deleted, created_at, updated_at
`
	sqlUserAdminDelete string = `
delete
from
    users
where
    id = $1
`
)

type UserAdminPgRepository struct {
	*repository.BaseRepository[*domain.User, string]
}

func NewUserAdminPgRepository(executor db.Executor, decipher db.ErrorDecipher) (*UserAdminPgRepository, error) {
	res := &UserAdminPgRepository{}
	// sql builders
	queryBuilders := repository.NewBaseQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlUserAdminFind
		}).
		WithList(func() string {
			return sqlUserAdminList
		}).
		WithCreate(func() string {
			return sqlUserAdminCreate
		}).
		WithChange(func() string {
			return sqlUserAdminChange
		}).
		WithDelete(func() string {
			return sqlUserAdminDelete
		}).
		Build()
	// callbacks
	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.User, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyUser).
		WithValidateCreate(res.validateCreate).
		WithBeforeCreate(res.beforeCreate).
		WithCreator(res.creator).
		WithValidateChange(res.validateChange).
		WithBeforeChange(res.beforeChange).
		WithChanger(res.changer).
		Build()
	// base CRUD
	base, err := repository.NewBaseRepository[*domain.User, string](
		executor,
		decipher,
		repository.NewEntityInfo("users", "User"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create UserPgRepository", err)
	}

	res.BaseRepository = base

	return res, nil
}

func (uar *UserAdminPgRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	querier := uar.GetExecutor().GetQuerier(ctx)

	row := querier.QueryRowContext(ctx, sqlUserAdminFindByName, name)

	res := uar.GetCallbacks().NewEntityFactory()

	err := uar.GetCallbacks().EntityScanner(row, res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewDalNotFoundError(uar.GetInfo().Entity, name, err)
		}

		return nil, errs.NewDalError("UserAdminPgRepository.FindByName", "get row", err)
	}

	if uar.GetCallbacks().AfterFind != nil {
		return uar.GetCallbacks().AfterFind(res)
	}

	return res, nil
}

func (uar *UserAdminPgRepository) entityScanner(scanner repository.Scannable, dest *domain.User) error {
	return scanner.Scan(&dest.ID, dest.Name, &dest.PasswordHash, &dest.Active, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (uar *UserAdminPgRepository) validateCreate(entity *domain.User) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateCreate()
}

func (uar *UserAdminPgRepository) beforeCreate(entity *domain.User) error {
	if err := entity.ValidateCreate(); err != nil {
		return errs.NewDalError("UserAdminPgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

func (uar *UserAdminPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.User) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, uar.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.PasswordHash, entity.CreatedAt, entity.UpdatedAt), nil
}

func (uar *UserAdminPgRepository) validateChange(entity *domain.User) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateChange()
}

func (uar *UserAdminPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.User) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, uar.GetQueryBuilders().GetChange()(), entity.ID, entity.PasswordHash, entity.Active, entity.Deleted, entity.UpdatedAt), nil
}

func (uar *UserAdminPgRepository) beforeChange(entity *domain.User) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("UserAdminPgRepository.beforeChange", "before change entity", err)
	}

	return nil
}

func (uar *UserAdminPgRepository) Close() error {
	return nil
}
