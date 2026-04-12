package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type RolePgRepository struct {
	*repository.BaseCRUDRepository[*domain.Role, string]
}

var _ libdomain.CRUDRepository[*domain.Role, string] = (*RolePgRepository)(nil)
var _ domain.RoleRepository = (*RolePgRepository)(nil)

//goland:noinspection DuplicatedCode
func NewRolePgRepository(executor db.Executor, decipher db.ErrorDecipher) (*RolePgRepository, error) {
	// new instance
	res := &RolePgRepository{}
	// sql builders
	queryBuilders := repository.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
	callbacks, _ := repository.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
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

func (rr *RolePgRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	if name == "" {
		return nil, errs.NewInvalidArgumentError("name", "is required")
	}

	return rr.GetHelper().Get(ctx, sqlRoleFindByName, name)
}

func (rr *RolePgRepository) entityScanner(scanner repository.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (rr *RolePgRepository) afterFind(entity *domain.Role, params ...any) (*domain.Role, error) {
	if entity.IsDeleted() {
		return nil, errs.NewDalSoftDeletedError(rr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, nil
}

func (rr *RolePgRepository) afterListYield(entity *domain.Role, params ...any) (*domain.Role, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(rr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

func (rr *RolePgRepository) validateCreate(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateCreate()
}

func (rr *RolePgRepository) beforeCreate(entity *domain.Role, params ...any) error {
	if err := entity.ValidateCreate(); err != nil {
		return errs.NewDalError("RolePgRepository.beforeCreate", "before create entity", err)
	}

	return nil
}

func (rr *RolePgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rr.GetQueryBuilders().GetCreate()(), entity.ID, entity.Name, entity.Description, entity.CreatedAt, entity.UpdatedAt), nil
}

func (rr *RolePgRepository) validateChange(entity *domain.Role, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "role entity is nil")
	}

	return entity.ValidateChange()
}

func (rr *RolePgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, rr.GetQueryBuilders().GetChange()(), entity.ID, entity.Name, entity.Description, entity.UpdatedAt), nil
}

func (rr *RolePgRepository) beforeChange(entity *domain.Role, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("RolePgRepository.beforeChange", "before change entity", err)
	}

	return nil
}
