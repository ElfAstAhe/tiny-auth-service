package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesAdminPgRepository struct {
	*repository.BaseOwnedRepository[*domain.Role, string, string]
}

var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesPgRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesAdminPgRepository)(nil)

func NewUserRolesAdminPgRepository(exec db.Executor, errDecipher db.ErrorDecipher) (*UserRolesAdminPgRepository, error) {
	res := &UserRolesAdminPgRepository{}
	// query builders
	queryBuilders := repository.NewBaseOwnedQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlUserRolesAdminFind
		}).
		WithListAll(func() string {
			return sqlUserRolesAdminListAll
		}).
		WithListAllByOwners(func() string {
			return sqlUserRolesAdminListAllByOwners
		}).
		WithCreate(func() string {
			return sqlUserRolesAdminCreate
		}).
		WithDelete(func() string {
			return sqlUserRolesAdminDelete
		}).
		WithDeleteAll(func() string {
			return sqlUserRolesAdminDeleteAll
		}).
		Build()

	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.Role, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyRole).
		WithValidateCreate(res.validateCreate).
		WithCreator(res.creator).
		Build()
	if err != nil {
		return nil, err
	}

	base, err := repository.NewBaseOwnedRepository[*domain.Role, string, string](
		exec,
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

func (ura *UserRolesAdminPgRepository) entityScanner(scanner repository.Scannable, sourceLabel string, dest *domain.Role, params ...any) error {
	switch sourceLabel {
	case repository.SourceLabelCreate:
		// ВНИМАНИЕ!!! Scan нужен, ибо без него получим driver: bad connection
		return scanner.Scan(&dest.ID)
	case repository.SourceLabelFind:
		return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	case repository.SourceLabelListAll:
		return scanner.Scan(&dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	case repository.SourceLabelListAllByOwners:
		return scanner.Scan(params[0], &dest.ID, &dest.Name, &dest.Description, &dest.Deleted, &dest.CreatedAt, &dest.UpdatedAt)
	}

	return errs.NewDalError("UserRolesAdminPgRepository.entityScanner", fmt.Sprintf("unknown source label [%v]", sourceLabel), nil)
}

func (ura *UserRolesAdminPgRepository) validateCreate(role *domain.Role, params ...any) error {
	if role == nil {
		return errs.NewInvalidArgumentError("role", "role is nil")
	}
	if strings.TrimSpace(role.ID) == "" {
		return errs.NewInvalidArgumentError("role", "role id is empty")
	}
	if len(params) == 0 {
		return errs.NewInvalidArgumentError("params", "params is empty")
	}
	userID, ok := params[0].(string)
	if !ok {
		return errs.NewInvalidArgumentError("userID", "must be a string")
	}
	if strings.TrimSpace(userID) == "" {
		return errs.NewInvalidArgumentError("userID", "must not be empty")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Role, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, ura.GetQueryBuilders().GetCreate()(), params[0], entity.ID), nil
}

func (ura *UserRolesAdminPgRepository) ValidateDeleteAll(ownerID string) error {
	if strings.TrimSpace(ownerID) == "" {
		return errs.NewInvalidArgumentError("ownerID", "must not be empty")
	}

	return nil
}

func (ura *UserRolesAdminPgRepository) ValidateDelete(ownerID string) error {
	if strings.TrimSpace(ownerID) == "" {
		return errs.NewInvalidArgumentError("ownerID", "must not be empty")
	}

	return nil
}
