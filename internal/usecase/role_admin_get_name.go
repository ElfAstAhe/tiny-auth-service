package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type RoleAdminGetNameUseCase interface {
	Get(ctx context.Context, name string) (*domain.Role, error)
}

type RoleAdminGetNameInteractor struct {
	roleRepo domain.RoleAdminRepository
}

func NewRoleAdminGetNameUseCase(roleRepo domain.RoleAdminRepository) *RoleAdminGetNameInteractor {
	return &RoleAdminGetNameInteractor{
		roleRepo: roleRepo,
	}
}

func (agn *RoleAdminGetNameInteractor) Get(ctx context.Context, name string) (*domain.Role, error) {
	if err := agn.validate(name); err != nil {
		return nil, domerrs.NewBllValidateError("RoleAdminGetNameInteractor.Get", "validate income data failed", err)
	}

	res, err := agn.roleRepo.FindByName(ctx, name)
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return nil, domerrs.NewBllNotFoundError("RoleAdminGetNameInteractor.Get", "Role", name, err)
		}

		return nil, domerrs.NewBllError("RoleAdminGetNameInteractor.Get", fmt.Sprintf("find Role model name [%s] failed", name), err)
	}

	return res, nil
}

func (agn *RoleAdminGetNameInteractor) validate(name string) error {
	if name == "" {
		return errs.NewInvalidArgumentError("name", "name is empty")
	}

	return nil
}
