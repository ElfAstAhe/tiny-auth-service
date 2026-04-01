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
	sqlUserAdminFind string = `
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
	sqlUserAdminFindByName string = `
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
	sqlUserAdminList string = `
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
	sqlUserAdminCreate string = `
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
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
`
	sqlUserAdminChange string = `
update
    users
set
    name = $2,
    user_type = $3,
    password_hash = $4,
    public_key = $5,
    private_key = $6,
    active = $7,
    deleted = $8,
    updated_at = $9
where
    id = $1
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
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
	*librepository.BaseCRUDRepository[*domain.User, string]
	userRolesRepo domain.UserRolesAdminRepository
	cipherHelper  helper.Cipher
	hashCipher    utils.Cipher
}

var _ domain.UserAdminRepository = (*UserAdminPgRepository)(nil)

func NewUserAdminPgRepository(executor db.Executor, errDecipher db.ErrorDecipher, cipherHelper helper.Cipher, hashCipher utils.Cipher, userRolesRepo domain.UserRolesAdminRepository) (*UserAdminPgRepository, error) {
	res := &UserAdminPgRepository{
		userRolesRepo: userRolesRepo,
		cipherHelper:  cipherHelper,
		hashCipher:    hashCipher,
	}
	// sql builders
	queryBuilders := librepository.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
		errDecipher,
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

func (uar *UserAdminPgRepository) Find(ctx context.Context, id string) (*domain.User, error) {
	res, err := uar.BaseCRUDRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	res.Roles, err = uar.userRolesRepo.ListAll(ctx, res.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uar *UserAdminPgRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	if name == "" {
		return nil, errs.NewInvalidArgumentError("name", "name is empty")
	}
	res, err := uar.GetHelper().Get(ctx, repository.SourceLabelFindByName, sqlUserAdminFindByName, name)
	if err != nil {
		return nil, err
	}
	roles, err := uar.userRolesRepo.ListAll(ctx, res.ID)
	if err != nil {
		return nil, err
	}
	res.Roles = roles

	return res, nil
}

func (uar *UserAdminPgRepository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	res, err := uar.BaseCRUDRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	allRoles, err := uar.userRolesRepo.ListAllByOwners(ctx, libdomain.EntitiesToIDList(res)...)
	if err != nil {
		return nil, err
	}

	for _, user := range res {
		if roles, ok := allRoles[user.ID]; ok {
			user.Roles = roles
		}
	}

	return res, nil
}

func (uar *UserAdminPgRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	res, err := uar.BaseCRUDRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	// сохраняем все привязки ролей
	roles, err := uar.userRolesRepo.Save(ctx, res.GetID(), user.Roles)
	if err != nil {
		return nil, err
	}
	res.Roles = roles

	return res, nil
}

func (uar *UserAdminPgRepository) Change(ctx context.Context, user *domain.User) (*domain.User, error) {
	res, err := uar.BaseCRUDRepository.Change(ctx, user)
	if err != nil {
		return nil, err
	}
	// сохраняем все привязки ролей
	roles, err := uar.userRolesRepo.Save(ctx, res.GetID(), user.Roles)
	if err != nil {
		return nil, err
	}
	res.Roles = roles

	return res, nil
}

func (uar *UserAdminPgRepository) Delete(ctx context.Context, id string) error {
	// удаляем привязки ролей
	err := uar.userRolesRepo.DeleteAll(ctx, id)
	if err != nil {
		return err
	}
	err = uar.BaseCRUDRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uar *UserAdminPgRepository) afterFind(entity *domain.User, params ...any) (*domain.User, error) {
	entity.PublicKey = uar.cipherHelper.DecryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.DecryptString(entity.PrivateKey)

	return entity, nil
}

func (uar *UserAdminPgRepository) afterListYield(entity *domain.User, params ...any) (*domain.User, bool, error) {
	entity.PublicKey = uar.cipherHelper.DecryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.DecryptString(entity.PrivateKey)

	return entity, true, nil
}

func (uar *UserAdminPgRepository) entityScanner(scanner librepository.Scannable, sourceLabel string, dest *domain.User, params ...any) error {
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

func (uar *UserAdminPgRepository) validateCreate(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateCreate()
}

func (uar *UserAdminPgRepository) beforeCreate(entity *domain.User, params ...any) error {
	if err := entity.BeforeCreate(); err != nil {
		return errs.NewDalError("UserAdminPgRepository.beforeCreate", "before create entity", err)
	}
	var err error
	entity.PublicKey = uar.cipherHelper.EncryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.EncryptString(entity.PrivateKey)
	entity.PasswordHash, err = uar.hashCipher.EncryptString(entity.PasswordHash)
	if err != nil {
		return errs.NewDalError("UserAdminPgRepository.beforeCreate", "encrypt (hash) password", err)
	}

	return nil
}

func (uar *UserAdminPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.User, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, uar.GetQueryBuilders().GetCreate()(),
		entity.ID,
		entity.Name,
		entity.Type,
		entity.PasswordHash,
		entity.PublicKey,
		entity.PrivateKey,
		entity.Active,
		entity.Deleted,
		entity.CreatedAt,
		entity.UpdatedAt,
	), nil
}

func (uar *UserAdminPgRepository) validateChange(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateChange()
}

func (uar *UserAdminPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.User, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, uar.GetQueryBuilders().GetChange()(),
		entity.ID,
		entity.Name,
		entity.Type,
		entity.PasswordHash,
		entity.PublicKey,
		entity.PrivateKey,
		entity.Active,
		entity.Deleted,
		entity.UpdatedAt,
	), nil
}

func (uar *UserAdminPgRepository) beforeChange(entity *domain.User, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("UserAdminPgRepository.beforeChange", "before change entity", err)
	}
	entity.PublicKey = uar.cipherHelper.EncryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.EncryptString(entity.PrivateKey)

	return nil
}
