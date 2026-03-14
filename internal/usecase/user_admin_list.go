package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type RoleAdminListUseCase interface {
	List(ctx context.Context, limit, offset int) ([]*domain.Role, error)
}

type RoleAdminListInteractor struct {
	roleRepo domain.RoleAdminRepository
}

func NewRoleAdminListUseCase(roleRepo domain.RoleAdminRepository) *RoleAdminListInteractor {
	return &RoleAdminListInteractor{roleRepo: roleRepo}
}

func (ral *RoleAdminListInteractor) List(ctx context.Context, limit, offset int) ([]*domain.Role, error) {
	res, err := ral.roleRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, errs.NewBllError("", fmt.Sprintf("list Role data with limit [%v] and offset [%v] failed", limit, offset), err)
	}

	return res, nil
}
