package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type RoleAdminListUseCase interface {
	List(ctx context.Context, limit, offset int) ([]*domain.Role, error)
}

type RoleAdminListInteractor struct {
	roleRepo     domain.RoleAdminRepository
	maxListLimit int
}

func NewRoleAdminListUseCase(roleRepo domain.RoleAdminRepository) *RoleAdminListInteractor {
	return &RoleAdminListInteractor{roleRepo: roleRepo}
}

func (ral *RoleAdminListInteractor) List(ctx context.Context, limit, offset int) ([]*domain.Role, error) {
	if err := ral.validate(limit, offset); err != nil {
		return nil, domerrs.NewBllValidateError("RoleAdminListInteractor.List", "validate income data failed", err)
	}

	res, err := ral.roleRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, domerrs.NewBllError("", fmt.Sprintf("list Role data with limit [%v] and offset [%v] failed", limit, offset), err)
	}

	return res, nil
}

func (ral *RoleAdminListInteractor) validate(limit, offset int) error {
	// correct limit
	if limit <= 0 {
		return errs.NewInvalidArgumentError("limit", "must be greater than zero")
	}
	// max limit
	if limit > ral.maxListLimit {
		return errs.NewInvalidArgumentError("limit", fmt.Sprintf("must be less than or equal to max limit [%v]", ral.maxListLimit))
	}
	// offset
	if offset < 0 {
		return errs.NewInvalidArgumentError("offset", "must be greater or equal than zero")
	}

	return nil
}
