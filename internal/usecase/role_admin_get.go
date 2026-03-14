package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type RoleAdminGetUseCase interface {
	Get(ctx context.Context, ID string) (*domain.Role, error)
}

type RoleAdminGetInteractor struct {
	roleRepo domain.RoleAdminRepository
}

func NewRoleAdminGetUseCase(roleRepo domain.RoleAdminRepository) *RoleAdminGetInteractor {
	return &RoleAdminGetInteractor{
		roleRepo: roleRepo,
	}
}

func (rag *RoleAdminGetInteractor) Get(ctx context.Context, ID string) (*domain.Role, error) {
	res, err := rag.roleRepo.Find(ctx, ID)
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return nil, domerrs.NewBllNotFoundError("RoleAdminGetInteractor.Get", "Role", ID, err)
		}

		return nil, domerrs.NewBllError("RoleAdminGetInteractor.Get", fmt.Sprintf("find role model id [%s] failed", ID), err)
	}

	return res, nil
}
