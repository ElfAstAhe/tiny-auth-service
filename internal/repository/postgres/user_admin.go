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

// UserAdminPgRepository structures administrative relational mapping adapters for comprehensive core user entity CRUD operations,
// orchestrating automated cryptographic credentials de-serialization hook events and batching cascaded RBAC owned roles persistence.
type UserAdminPgRepository struct {
	*librepo.BaseCRUDRepository[*domain.User, string]                                 // Generic platform core structural database repository handle
	userRolesRepo                                     domain.UserRolesAdminRepository // Downstream administrative sub-repository handle managing relational bridge role mappings
	cipherHelper                                      helper.Cipher                   // Structural symmetric encryption utility decrypting persisted cryptographic key blocks
	hashCipher                                        utils.Cipher                    // Cryptographic tool managing comparative password hashing mechanics
}

// Compile-time interface compliance verifications
var _ libdom.CRUDRepository[*domain.User, string] = (*UserAdminPgRepository)(nil)
var _ domain.UserAdminRepository = (*UserAdminPgRepository)(nil)

// NewUserAdminPgRepository acts as a factory constructor compiling query configurations, scanner mappings, cascaded lifecycle hooks, and validation criteria.
func NewUserAdminPgRepository(
	executor db.Executor,
	errDecipher db.ErrorDecipher,
	cipherHelper helper.Cipher,
	hashCipher utils.Cipher,
	userRolesRepo domain.UserRolesAdminRepository,
) (*UserAdminPgRepository, error) {
	res := &UserAdminPgRepository{
		userRolesRepo: userRolesRepo,
		cipherHelper:  cipherHelper,
		hashCipher:    hashCipher,
	}
	// sql builders
	queryBuilders := librepo.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
		errDecipher,
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

// Find retrieves a single user aggregate record by its specific identifier, performing cascaded eager loading to append associated RBAC roles.
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

// FindByName executes a specialized index-based user aggregate retrieval sequence matching target unique string name credentials parameters.
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

// List fetches a paginated array collection of user models, executing optimized single-query batch extraction for all related user role aggregates to mitigate N+1 pitfalls.
func (uar *UserAdminPgRepository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	res, err := uar.BaseCRUDRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	allRoles, err := uar.userRolesRepo.ListAllByOwners(ctx, libdom.EntitiesToIDList(res)...)
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

// Create persists a new core user record structure and triggers transactional cascade insertions saving attached role assignment definitions.
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

// Change updates an existing core user record mapping and forces synchronized structural replacements updating the relational role bridge mapping.
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

// Delete manages cascaded administrative cleanups, enforcing strict database relational clearing of bridge role assets prior to dropping the user entry.
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

// afterFind handles post-retrieval pipeline event routines to decrypt structural user cryptographic key pairs on individual entity records.
func (uar *UserAdminPgRepository) afterFind(entity *domain.User, params ...any) (*domain.User, error) {
	entity.PublicKey = uar.cipherHelper.DecryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.DecryptString(entity.PrivateKey)

	return entity, nil
}

// afterListYield performs post-retrieval stream mutations decrypting sensitive key components across collection iterator batches.
func (uar *UserAdminPgRepository) afterListYield(entity *domain.User, params ...any) (*domain.User, bool, error) {
	entity.PublicKey = uar.cipherHelper.DecryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.DecryptString(entity.PrivateKey)

	return entity, true, nil
}

// entityScanner maps raw relational SQL row column values directly into a concrete user aggregate memory model pointer.
func (uar *UserAdminPgRepository) entityScanner(scanner librepo.Scannable, sourceLabel string, dest *domain.User, params ...any) error {
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

// validateCreate evaluates structural model domain assertions before routing executions to insertion statements.
func (uar *UserAdminPgRepository) validateCreate(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateCreate()
}

// beforeCreate intercept execution sequences to run domain hooks and symmetrically encrypt structural key blocks or hash raw passwords before СУБД streaming.
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

// creator triggers low-level context execution routines injecting the fully encrypted user identity aggregate record into PostgreSQL.
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

// validateChange evaluates state invariants on user domain models prior to sending modification commands down the line.
func (uar *UserAdminPgRepository) validateChange(entity *domain.User, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "user entity is nil")
	}

	return entity.ValidateChange()
}

// changer maps modified memory fields context parameter boundaries back onto administrative database columns.
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

// beforeChange runs pre-modification entity workflows and forces secure key transformations before updating persistence entries.
func (uar *UserAdminPgRepository) beforeChange(entity *domain.User, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("UserAdminPgRepository.beforeChange", "before change entity", err)
	}
	entity.PublicKey = uar.cipherHelper.EncryptString(entity.PublicKey)
	entity.PrivateKey = uar.cipherHelper.EncryptString(entity.PrivateKey)

	return nil
}
