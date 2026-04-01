package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	librepository "github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository"
)

const (
	sqlUserFind string = `
select
    id,
    name,
    user_type,
    password_hash,
    public_key,
    private_key,
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
    user_type,
    password_hash,
    public_key,
    private_key,
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
    user_type,
    password_hash,
    public_key,
    private_key,
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
    user_type,
    password_hash,
    public_key,
    private_key,
    active,
    deleted,
    created_at,
    updated_at
)
values($1, $2, $3, $4, $5, $6, $7, false, $8, $9)
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
`
	sqlUserChange string = `
update
    users
set
    user_type = $2,
    password_hash = $3,
    public_key = $4,
    private_key = $5,
    active = $6,
    deleted = $7,
    updated_at = $8
where
    id = $1
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
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
	*librepository.BaseCRUDRepository[*domain.User, string]
	hashCipher    utils.Cipher
	cipherHelper  helper.Cipher
	userRolesRepo domain.UserRolesRepository
}

var _ domain.UserRepository = (*UserPgRepository)(nil)

//goland:noinspection DuplicatedCode
func NewUserPgRepository(executor db.Executor, decipher db.ErrorDecipher, hashCipher utils.Cipher, cipherHelper helper.Cipher, userRolesRepo domain.UserRolesRepository) (*UserPgRepository, error) {
	res := &UserPgRepository{
		hashCipher:    hashCipher,
		cipherHelper:  cipherHelper,
		userRolesRepo: userRolesRepo,
	}
	// sql builders
	queryBuilders := librepository.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
	callbacks, err := librepository.NewBaseRepositoryCallbacksBuilder[*domain.User, string]().NewInstance().
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
	base, err := librepository.NewBaseCRUDRepository[*domain.User, string](
		executor,
		decipher,
		librepository.NewEntityInfo("users", "User"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create UserPgRepository", err)
	}

	res.BaseCRUDRepository = base

	return res, nil
}

func (ur *UserPgRepository) Find(ctx context.Context, id string) (*domain.User, error) {
	res, err := ur.BaseCRUDRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	roles, err := ur.userRolesRepo.ListAll(ctx, id)
	if err != nil {
		return nil, err
	}
	res.Roles = roles

	return res, nil
}

func (ur *UserPgRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	if name == "" {
		return nil, errs.NewInvalidArgumentError("name", "name is empty")
	}
	res, err := ur.GetHelper().Get(ctx, repository.SourceLabelFindByName, sqlUserFindByName, name)
	if err != nil {
		return nil, err
	}
	roles, err := ur.userRolesRepo.ListAll(ctx, res.GetID())
	if err != nil {
		return nil, err
	}
	res.Roles = roles

	return res, nil
}

func (ur *UserPgRepository) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	res, err := ur.BaseCRUDRepository.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	// получаем списки ролей в разрезе UserID
	allRoles, err := ur.userRolesRepo.ListAllByOwners(ctx, libdomain.EntitiesToIDList(res)...)
	if err != nil {
		return nil, err
	}

	for _, user := range res {
		if roles, ok := allRoles[user.GetID()]; ok {
			user.Roles = roles
		}
	}

	return res, nil
}

func (ur *UserPgRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	res, err := ur.BaseCRUDRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	// сохраняем все привязки ролей
	roles, err := ur.userRolesRepo.Save(ctx, res.GetID(), user.Roles)
	if err != nil {
		return nil, err
	}
	res.Roles = roles

	return res, nil
}

func (ur *UserPgRepository) entityScanner(scanner librepository.Scannable, sourceLabel string, dest *domain.User, params ...any) error {
	return scanner.Scan(
		&dest.ID,
		&dest.Name,
		&dest.Type,
		&dest.PasswordHash,
		&dest.PublicKey,
		&dest.PrivateKey,
		&dest.Active,
		&dest.Deleted,
		&dest.CreatedAt,
		&dest.UpdatedAt,
	)
}

func (ur *UserPgRepository) afterFind(entity *domain.User, params ...any) (*domain.User, error) {
	if entity.IsDeleted() {
		return nil, errs.NewDalSoftDeletedError(ur.GetHelper().GetInfo().Entity, entity.GetID())
	}

	entity.PublicKey = ur.cipherHelper.DecryptString(entity.PublicKey)
	entity.PrivateKey = ur.cipherHelper.DecryptString(entity.PrivateKey)

	return entity, nil
}

func (ur *UserPgRepository) afterListYield(entity *domain.User, params ...any) (*domain.User, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(ur.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

func (ur *UserPgRepository) validateCreate(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateCreate()
}

func (ur *UserPgRepository) beforeCreate(entity *domain.User, params ...any) error {
	if err := entity.BeforeCreate(); err != nil {
		return errs.NewDalError("UserPgRepository.beforeCreate", "before create entity", err)
	}
	var err error
	entity.PublicKey = ur.cipherHelper.EncryptString(entity.PublicKey)
	entity.PrivateKey = ur.cipherHelper.EncryptString(entity.PrivateKey)
	entity.PasswordHash, err = ur.hashCipher.EncryptString(entity.PasswordHash)
	if err != nil {
		return errs.NewDalError("UserPgRepository.beforeCreate", "encrypt (hash) password", err)
	}

	return nil
}

func (ur *UserPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.User, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, ur.GetQueryBuilders().GetCreate()(),
		entity.ID,
		entity.Name,
		entity.Type,
		entity.PasswordHash,
		entity.PublicKey,
		entity.PrivateKey,
		entity.Active,
		entity.CreatedAt,
		entity.UpdatedAt,
	), nil
}

func (ur *UserPgRepository) validateChange(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateChange()
}

func (ur *UserPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.User, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, ur.GetQueryBuilders().GetChange()(),
		entity.ID,
		entity.Type,
		entity.PasswordHash,
		entity.PublicKey,
		entity.PrivateKey,
		entity.Active,
		entity.Deleted,
		entity.UpdatedAt,
	), nil
}

func (ur *UserPgRepository) beforeChange(entity *domain.User, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("UserPgRepository.beforeChange", "before change entity", err)
	}
	entity.PublicKey = ur.cipherHelper.EncryptString(entity.PublicKey)
	entity.PrivateKey = ur.cipherHelper.EncryptString(entity.PrivateKey)

	return nil
}
