package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	librepo "github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository"
)

// RoleAdminPgRepository structures administrative relational mapping adapters for generic core role entity CRUD operations.
type RoleAdminPgRepository struct {
	*librepo.BaseCRUDRepository[*domain.Role, string] // Generic platform core structural database repository handle
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleAdminPgRepository)(nil)
var _ domain.RoleAdminRepository = (*RoleAdminPgRepository)(nil)

// NewRoleAdminPgRepository acts as a factory constructor compiling query configurations, scanner mappings, and lifecycle validations criteria.
func NewRoleAdminPgRepository(
	executor db.Executor,
	decipher db.ErrorDecipher,
) (*RoleAdminPgRepository, error) {
	// new instance
	res := &RoleAdminPgRepository{}
	// sql builders
	queryBuilders := librepo.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
	callbacks, _ := librepo.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
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
	base, err := librepo.NewBaseCRUDRepository[*domain.Role, string](
		executor,
		decipher,
		librepo.NewEntityInfo("roles", "Role"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create RolePgRepository", err)
	}

	res.BaseCRUDRepository = base

	return res, nil
}

// FindByName executes a fine-grained lookup operation matching the string role name utilizing low-level platform helper mechanisms.
func (rar *RoleAdminPgRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	if name == "" {
		return nil, errs.NewInvalidArgumentError("name", "cannot be empty")
	}

	return rar.GetHelper().Get(ctx, repository.SourceLabelFindByName, sqlRoleAdminFindByName, name)
}

// entityScanner maps raw relational row column values into concrete memory model pointers.
func (rar *RoleAdminPgRepository) entityScanner(scanner librepo.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

// validateCreate triggers domain aggregate assertions prior to executing SQL creation boundaries.
func (rar *RoleAdminPgRepository) validateCreate(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateCreate()
}

// beforeCreate delegates execution targets to model pre-persistence hook procedures.
func (rar *RoleAdminPgRepository) beforeCreate(entity *domain.Role, params ...any) error {
	if err := entity.BeforeCreate(); err != nil {
		return errs.NewDalError("RoleAdminPgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

// creator triggers context selections executing statement evaluations to insert a new role record into PostgreSQL.
func (rar *RoleAdminPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rar.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.Description, entity.CreatedAt, entity.UpdatedAt), nil
}

// validateChange triggers domain aggregate state verification rules prior to executing SQL modifications.
func (rar *RoleAdminPgRepository) validateChange(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateChange()
}

// changer executes context statements mapping changes back into administrative relational tables columns.
func (rar *RoleAdminPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rar.GetQueryBuilders().GetChange()(), entity.ID, entity.Name, entity.Description, entity.Deleted, entity.UpdatedAt), nil
}

// beforeChange delegates execution parameters to model pre-modification hook procedures.
func (rar *RoleAdminPgRepository) beforeChange(entity *domain.Role, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("RoleAdminPgRepository.beforeChange", "before change entity", err)
	}

	return nil
}
