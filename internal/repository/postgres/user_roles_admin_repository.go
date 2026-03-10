package postgres

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

const (
	sqlUserRolesAdminListByUser string = `
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
where
    ur.user_id = $1
`
	sqlUserRolesAdminListByUsers string = `
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
where
    ur.user_id = any($1)
`
	sqlUserRolesAdminCreate string = `
insert into user_roles (
    user_id,
    role_id
)
values ($1, $2)
`
	sqlUserRolesAdminDelete string = `
delete from user_roles where user_id = $1
`
)

type UserRolesAdminPgRepository struct {
	exec db.Executor
}

func NewUserRolesAdminPgRepository(exec db.Executor) *UserRolesAdminPgRepository {
	return &UserRolesAdminPgRepository{
		exec: exec,
	}
}

func (ura *UserRolesAdminPgRepository) ListByUser(ctx context.Context, userID string) ([]*domain.Role, error) {
	if err := ura.validateListByUser(userID); err != nil {
		return nil, errs.NewDalError("UserRolesPgRepository.ListByUser", "validate income params", err)
	}

	res, err := ura.internalList(ctx, sqlUserRolesListByUser, userID)
	if err != nil {
		return nil, errs.NewDalError("UserRolesPgRepository.ListByUser", "internal list", err)
	}

	return res[userID], nil
}

func (ura *UserRolesAdminPgRepository) validateListByUser(userID string) error {
	if userID == "" {
		return errs.NewInvalidArgumentError("userID", "is required")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) ListByUsers(ctx context.Context, userIDs ...string) (map[string][]*domain.Role, error) {
	if err := ura.validateListByUsers(userIDs...); err != nil {
		return nil, errs.NewDalError("UserRolesAdminPgRepository.ListByUsers", "validate income params", err)
	}

	return ura.internalList(ctx, sqlUserRolesAdminListByUsers, userIDs)
}

func (ura *UserRolesAdminPgRepository) validateListByUsers(userIDs ...string) error {
	if len(userIDs) == 0 {
		return errs.NewInvalidArgumentError("userIDs", "is required")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) Save(ctx context.Context, userID string, roles []*domain.Role) error {
	if err := ura.validateSave(userID, roles); err != nil {
		return err
	}

	if err := ura.Delete(ctx, userID); err != nil {
		return errs.NewDalError("UserRolesAdminPgRepository.Save", fmt.Sprintf("remove roles for user id [%s]", userID), err)
	}

	for _, role := range roles {
		if err := ura.create(ctx, userID, role); err != nil {
			return errs.NewDalError("UserRolesAdminPgRepository.Save", fmt.Sprintf("add role id [%s] for user id [%s]", role.ID, userID), err)
		}
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) validateSave(userID string, roles []*domain.Role) error {
	if userID == "" {
		return errs.NewInvalidArgumentError("userID", "is required")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) create(ctx context.Context, userID string, role *domain.Role) error {
	querier := ura.exec.GetQuerier(ctx)

	_, err := querier.ExecContext(ctx, sqlUserRolesAdminCreate, userID, role.ID)
	if err != nil {
		return errs.NewDalError("UserRolesAdminPgRepository.create", "exec sql", err)
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) Delete(ctx context.Context, userID string) error {
	if err := ura.validateDelete(userID); err != nil {
		return err
	}

	querier := ura.exec.GetQuerier(ctx)

	_, err := querier.ExecContext(ctx, sqlUserRolesAdminDelete, userID)
	if err != nil {
		return errs.NewDalError("UserRolesAdminPgRepository.Delete", "exec sql", err)
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) validateDelete(userID string) error {
	if userID == "" {
		return errs.NewInvalidArgumentError("userID", "is required")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) internalList(ctx context.Context, sqlReq string, params ...any) (map[string][]*domain.Role, error) {
	querier := ura.exec.GetQuerier(ctx)

	rows, err := querier.QueryContext(ctx, sqlReq, params...)
	if err != nil {
		return nil, errs.NewDalError("UserRolesPgRepository.internalList", "query", err)
	}
	defer rows.Close()

	res := make(map[string][]*domain.Role)
	for rows.Next() {
		if err = ctx.Err(); err != nil {
			return nil, errs.NewDalError("UserRolesPgRepository.internalList", "check context", err)
		}

		entity := domain.NewEmptyRole()
		userID := ""

		err = ura.scan(rows, userID, entity)
		if err != nil {
			return nil, errs.NewDalError("UserRolesPgRepository.internalList", "scan rows", err)
		}
		if _, ok := res[userID]; !ok {
			res[userID] = make([]*domain.Role, 0)
		}

		res[userID] = append(res[userID], entity)
	}
	if rows.Err() != nil {
		return nil, errs.NewDalError("UserRolesPgRepository.internalList", "after scan", rows.Err())
	}

	return res, nil
}

func (ura *UserRolesAdminPgRepository) scan(scanner repository.Scannable, userID string, dest *domain.Role) error {
	return scanner.Scan(&userID, &dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
}
