package postgres

import (
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
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
	sqlUserRolesDeleteAll string = `
delete from
    user_roles
where
    user_id = $1
`
	sqlUserRolesCreate string = `
insert into user_roles(
    user_id,
    role_id
)
values ($1, $2)
returning role_id
`
)

type UserRolesPgRepository struct {
	*repository.BaseOwnedRepository[*domain.Role, string, string]
}

var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesPgRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesPgRepository)(nil)

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

func (urr *UserRolesPgRepository) entityScanner(scanner repository.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	switch sourceLabel {
	case repository.SourceLabelListAll:
		return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	case repository.SourceLabelListAllByOwners:
		return scanner.Scan(params[0], &dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	}

	return errs.NewDalError("UserRolesPgRepository.entityScanner", fmt.Sprintf("unknown source label [%v]", sourceLabel), nil)
}

func (urr *UserRolesPgRepository) afterListYield(entity *domain.Role, params ...any) (*domain.Role, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(urr.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}
