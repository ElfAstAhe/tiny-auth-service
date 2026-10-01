package postgres

import (
	"context"
	"database/sql"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	librepo "github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository"
)

// RolePgRepository structures concrete relational mapping adapters for generic core role entity CRUD operations,
// implementing lifecycle soft-delete validation filters across active single and collection record streams.
type RolePgRepository struct {
	*librepo.BaseCRUDRepository[*domain.Role, string] // Generic platform core structural database repository handle
}

// Compile-time interface compliance verifications
var _ libdom.CRUDRepository[*domain.Role, string] = (*RolePgRepository)(nil)
var _ domain.RoleRepository = (*RolePgRepository)(nil)

// NewRolePgRepository acts as a factory constructor compiling query configurations, scanner mappings, and cascading callback pipelines.
//
//goland:noinspection DuplicatedCode
func NewRolePgRepository(
	executor db.Executor,
	decipher db.ErrorDecipher,
) (*RolePgRepository, error) {
	// new instance
	res := &RolePgRepository{}
	// sql builders
	queryBuilders := librepo.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
	callbacks, _ := librepo.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
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

// FindByName executes a fine-grained lookup operation matching target unique string name parameters utilizing low-level platform helper mechanisms.
func (rr *RolePgRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errs.NewInvalidArgumentError("name", "is required")
	}

	return rr.GetHelper().Get(ctx, repository.SourceLabelFindByName, sqlRoleFindByName, name)
}

// entityScanner maps raw relational SQL row column values directly into a concrete role domain model pointer.
func (rr *RolePgRepository) entityScanner(scanner librepo.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

// afterFind handles post-retrieval pipeline intercept routines to transparently filter out soft-deleted single roles records.
func (rr *RolePgRepository) afterFind(entity *domain.Role, params ...any) (*domain.Role, error) {
	if entity.IsDeleted() {
		return nil, errs.NewDalSoftDeletedError(rr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, nil
}

// afterListYield handles post-retrieval stream mutations to seamlessly drop and skip soft-deleted role data fields across collection chunks.
func (rr *RolePgRepository) afterListYield(entity *domain.Role, params ...any) (*domain.Role, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(rr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

// validateCreate evaluates core structural domain assertions before passing execution to the SQL insertion layer.
func (rr *RolePgRepository) validateCreate(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateCreate()
}

// beforeCreate triggers pre-persistence domain validations prior to formatting actual data insertions payloads.
func (rr *RolePgRepository) beforeCreate(entity *domain.Role, params ...any) error {
	if err := entity.ValidateCreate(); err != nil {
		return errs.NewDalError("RolePgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

// creator triggers low-level context execution routines writing a new role entity configuration record to PostgreSQL.
func (rr *RolePgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rr.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.Description, entity.CreatedAt, entity.UpdatedAt), nil
}

// validateChange asserts state criteria bounds on user-facing memory fields prior to committing database modifications.
func (rr *RolePgRepository) validateChange(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateChange()
}

// changer executes context expressions mapping current field data back into relational database columns.
func (rr *RolePgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rr.GetQueryBuilders().GetChange()(), entity.ID, entity.Name, entity.Description, entity.UpdatedAt), nil
}

// beforeChange maps execution blocks to pre-modification domain validation sequences.
func (rr *RolePgRepository) beforeChange(entity *domain.Role, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("RolePgRepository.beforeChange", "before change entity", err)
	}

	return nil
}
