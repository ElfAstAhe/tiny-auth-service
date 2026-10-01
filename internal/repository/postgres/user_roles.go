package postgres

import (
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRolesPgRepository structures public relational mapping adapters for many-to-many user roles associations.
type UserRolesPgRepository struct {
	*repository.BaseOwnedRepository[*domain.Role, string, string] // Generic platform core structural owned database repository handle
}

// Compile-time interface compliance verifications
var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesPgRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesPgRepository)(nil)

// NewUserRolesPgRepository acts as a factory constructor compiling query configurations, list lifecycles, and relational database footprints.
func NewUserRolesPgRepository(
	executor db.Executor,
	errDecipher db.ErrorDecipher,
) (*UserRolesPgRepository, error) {
	res := &UserRolesPgRepository{}

	// query builders
	queryBuilders := repository.NewBaseOwnedQueryBuildersBuilder().NewInstance().
		WithListAll(func() string {
			return sqlUserRolesListAll
		}).
		WithListAllByOwners(func() string {
			return sqlUserRolesListAllByOwners
		}).
		WithDeleteAll(func() string {
			return sqlUserRolesDeleteAll
		}).
		WithCreate(func() string {
			return sqlUserRolesCreate
		}).
		Build()

	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyRole).
		WithAfterListYield(res.afterListYield).
		Build()
	if err != nil {
		return nil, err
	}

	base, err := repository.NewBaseOwnedRepository[*domain.Role, string, string](
		executor,
		errDecipher,
		repository.NewEntityInfo("user_roles", "UserRole"),
		queryBuilders,
		callbacks,
		repository.LinkStrategyManyToMany,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res.BaseOwnedRepository = base

	return res, nil
}

// entityScanner maps raw relational row column values into concrete memory model pointers based on scanning label criteria.
func (urr *UserRolesPgRepository) entityScanner(scanner repository.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	switch sourceLabel {
	case repository.SourceLabelListAll:
		return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	case repository.SourceLabelListAllByOwners:
		return scanner.Scan(params[0], &dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	}

	return errs.NewDalError("UserRolesPgRepository.entityScanner", fmt.Sprintf("unknown source label [%v]", sourceLabel), nil)
}

// afterListYield runs a post-retrieval pipeline check to seamlessly isolate and drop softly-deleted entities from active collection streams.
func (urr *UserRolesPgRepository) afterListYield(entity *domain.Role, params ...any) (*domain.Role, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(urr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}
