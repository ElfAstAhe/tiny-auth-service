package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	librepo "github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository"
)

// UserPgRepository structures relational mapping adapters for generic public user entity CRUD operations,
// implementing transparent cryptographic serialization boundaries and eager-loading RBAC associations loops.
type UserPgRepository struct {
	*librepo.BaseCRUDRepository[*domain.User, string]                            // Generic platform core structural database repository handle
	hashCipher                                        utils.Cipher               // Cryptographic tool managing comparative password hashing mechanics
	cipherHelper                                      helper.Cipher              // Structural symmetric encryption utility decrypting persisted cryptographic key blocks
	userRolesRepo                                     domain.UserRolesRepository // Relational sub-repository handle managing public role bridge mappings
}

// Compile-time interface compliance verifications
var _ libdom.CRUDRepository[*domain.User, string] = (*UserPgRepository)(nil)
var _ domain.UserRepository = (*UserPgRepository)(nil)

// NewUserPgRepository acts as a factory constructor compiling query configurations, scanner mappings, cascading callback pipelines, and crypto engines.
func NewUserPgRepository(
	executor db.Executor,
	decipher db.ErrorDecipher,
	hashCipher utils.Cipher,
	cipherHelper helper.Cipher,
	userRolesRepo domain.UserRolesRepository,
) (*UserPgRepository, error) {
	res := &UserPgRepository{
		hashCipher:    hashCipher,
		cipherHelper:  cipherHelper,
		userRolesRepo: userRolesRepo,
	}
	// sql builders
	queryBuilders := librepo.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
	callbacks, _ := librepo.NewBaseRepositoryCallbacksBuilder[*domain.User, string]().NewInstance().
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
	base, err := librepo.NewBaseCRUDRepository[*domain.User, string](
		executor,
		decipher,
		librepo.NewEntityInfo("users", "User"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create UserPgRepository", err)
	}

	res.BaseCRUDRepository = base

	return res, nil
}

// Find retrieves a single user aggregate record by identifier, returning soft-delete evaluation metrics or loading attached roles.
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

// FindByName executes a customized lookup sequence matching unique string username credentials via platform helper mechanisms.
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

// List fetches a paginated array collection of user records, triggering optimized single-query batch extraction to link owned role maps.
func (ur *UserPgRepository) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	res, err := ur.BaseCRUDRepository.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	// получаем списки ролей в разрезе UserID
	allRoles, err := ur.userRolesRepo.ListAllByOwners(ctx, libdom.EntitiesToIDList(res)...)
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

// Create persists a new user record structure and maps associated role allocations inside a transactional cascade.
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

// entityScanner maps raw relational SQL row column fields directly into a concrete user memory model pointer.
func (ur *UserPgRepository) entityScanner(scanner librepo.Scannable, sourceLabel string, dest *domain.User, params ...any) error {
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

// afterFind handles post-retrieval pipeline intercept routines to transparently filter out soft-deleted records and decrypt key pair blocks.
func (ur *UserPgRepository) afterFind(entity *domain.User, params ...any) (*domain.User, error) {
	if entity.IsDeleted() {
		return nil, errs.NewDalSoftDeletedError(ur.GetHelper().GetInfo().Entity, entity.GetID())
	}

	entity.PublicKey = ur.cipherHelper.DecryptString(entity.PublicKey)
	entity.PrivateKey = ur.cipherHelper.DecryptString(entity.PrivateKey)

	return entity, nil
}

// afterListYield performs post-retrieval stream mutations isolating deleted records across collection chunks.
func (ur *UserPgRepository) afterListYield(entity *domain.User, params ...any) (*domain.User, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(ur.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

// validateCreate evaluates structural model domain assertions before passing execution to the SQL insertion layer.
func (ur *UserPgRepository) validateCreate(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateCreate()
}

// beforeCreate intercept execution sequences to run domain hooks and symmetrically encrypt structural key blocks or hash raw passwords before СУБД streaming.
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

// creator triggers low-level context execution routines injecting the fully encrypted user identity aggregate record into PostgreSQL.
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

// validateChange evaluates state invariants on user domain models prior to sending modification commands down the line.
func (ur *UserPgRepository) validateChange(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateChange()
}

// changer maps modified memory fields context parameter boundaries back onto relational database columns.
func (ur *UserPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.User, params ...any) (*sql.Row, error) {
	// ARGUMENT DRIFT NOTICE: Notice that entity.Name is omitted inside this specific parameter cascade sequence.
	// Ensure that sqlUserChange payload expectations are perfectly aligned to prevent driver interpolation failures.
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

// beforeChange runs pre-modification entity workflows and forces secure key transformations before updating persistence entries.
func (ur *UserPgRepository) beforeChange(entity *domain.User, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("UserPgRepository.beforeChange", "before change entity", err)
	}
	entity.PublicKey = ur.cipherHelper.EncryptString(entity.PublicKey)
	entity.PrivateKey = ur.cipherHelper.EncryptString(entity.PrivateKey)

	return nil
}
