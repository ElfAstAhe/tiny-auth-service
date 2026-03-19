package postgres

import (
	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

const (
	sqlUserRolesListAll string = `
select
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
        and r.deleted = false
where
    ur.user_id = $1
`
	sqlUserRolesListAllByOwners string = `
select
    ur.user_id,
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
        and r.deleted = false
where
    ur.user_id = any($1)
order by
    1 asc, 2 asc
`
)

type UserRolesPgRepository struct {
	*repository.BaseOwnedRepository[*domain.Role, string, string]
}

func NewUserRolesPgRepository(executor db.Executor, errDecipher db.ErrorDecipher) (*UserRolesPgRepository, error) {
	res := &UserRolesPgRepository{}

	// query builders
	queryBuilders := repository.NewBaseOwnedQueryBuildersBuilder().NewInstance().
		WithListAll(func() string {
			return sqlUserRolesListAll
		}).
		WithListAllByOwners(func() string {
			return sqlUserRolesListAllByOwners
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

func (urr *UserRolesPgRepository) entityScanner(scanner repository.Scannable, dest *domain.Role, params ...any) error {
	if len(params) == 0 {
		return scanner.Scan(&dest.ID, dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	}

	return scanner.Scan(&params[0], &dest.ID, dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}

func (urr *UserRolesPgRepository) afterListYield(entity *domain.Role, params ...any) (*domain.Role, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(urr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}
