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
	sqlUserFind string = `
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
	sqlUserFindByName string = `
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
	sqlUserList string = `
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
	sqlUserCreate string = `
insert into users (
    id,
    name,
    password_hash,
    active,
    deleted,
    created_at,
    updated_at
)
values($1, $2, $3, true, false, $4, $5)
returning id, name, password_hash, active, deleted, created_at, updated_at
`
	sqlUserChange string = `
update
    users
set
    password_hash = $2,
    active = $3,
    deleted = $4,
    updated_at = $5
where
    id = $1
returning id, name, password_hash, active, deleted, created_at, updated_at
`
	sqlUserDelete string = `
update
    users
set
    deleted = true
where
    id = $1
`
)

type UserPgRepository struct {
	*repository.BaseRepository[*domain.User, string]
}

func NewUserPgRepository(executor db.Executor, decipher db.ErrorDecipher) (*UserPgRepository, error) {
	res := &UserPgRepository{}
	// sql builders
	queryBuilders := repository.NewBaseQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlUserFind
		}).
		WithList(func() string {
			return sqlUserList
		}).
		WithCreate(func() string {
			return sqlUserCreate
		}).
		WithChange(func() string {
			return sqlUserChange
		}).
		WithDelete(func() string {
			return sqlUserDelete
		}).
		Build()
	// callbacks
	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.User, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyUser).
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

func (ur *UserPgRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	querier := ur.GetExecutor().GetQuerier(ctx)

	row := querier.QueryRowContext(ctx, sqlUserFindByName, name)

	res := ur.GetCallbacks().NewEntityFactory()

	err := ur.GetCallbacks().EntityScanner(row, res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewDalNotFoundError(ur.GetInfo().Entity, name, err)
		}

		return nil, errs.NewDalError("UserPgRepository.FindByName", "get row", err)
	}

	if ur.GetCallbacks().AfterFind != nil {
		return ur.GetCallbacks().AfterFind(res)
	}

	return res, nil
}

func (ur *UserPgRepository) entityScanner(scanner repository.Scannable, dest *domain.User) error {
	return scanner.Scan(&dest.ID, dest.Name, &dest.PasswordHash, &dest.Active, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (ur *UserPgRepository) afterFind(entity *domain.User) (*domain.User, error) {
	if entity.IsDeleted() {
		return nil, errs.NewDalSoftDeletedError(ur.GetInfo().Entity, entity.GetID())
	}

	return entity, nil
}

func (ur *UserPgRepository) afterListYield(entity *domain.User) (*domain.User, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(ur.GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

func (ur *UserPgRepository) validateCreate(entity *domain.User) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateCreate()
}

func (ur *UserPgRepository) beforeCreate(entity *domain.User) error {
	if err := entity.ValidateCreate(); err != nil {
		return errs.NewDalError("UserPgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

func (ur *UserPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.User) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, ur.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.PasswordHash, entity.CreatedAt, entity.UpdatedAt), nil
}

func (ur *UserPgRepository) validateChange(entity *domain.User) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateChange()
}

func (ur *UserPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.User) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, ur.GetQueryBuilders().GetChange()(), entity.ID, entity.PasswordHash, entity.Active, entity.Deleted, entity.UpdatedAt), nil
}

func (ur *UserPgRepository) beforeChange(entity *domain.User) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("UserPgRepository.beforeChange", "before change entity", err)
	}

	return nil
}

func (ur *UserPgRepository) Close() error {
	return nil
}
