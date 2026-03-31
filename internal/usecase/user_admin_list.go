package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type UserAdminListUseCase interface {
	List(ctx context.Context, limit, offset int) ([]*domain.User, error)
}

type UserAdminListInteractor struct {
	userRepo     domain.UserAdminRepository
	maxListLimit int
}

var _ UserAdminListUseCase = (*UserAdminListInteractor)(nil)

func NewUserAdminListUseCase(userRepo domain.UserAdminRepository, maxListLimit int) *UserAdminListInteractor {
	res := &UserAdminListInteractor{
		userRepo:     userRepo,
		maxListLimit: maxListLimit,
	}
	if res.maxListLimit < 0 {
		res.maxListLimit = DefaultMaxLimit
	}

	return res
}

func (ual *UserAdminListInteractor) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	if err := ual.validate(limit, offset); err != nil {
		return nil, domerrs.NewBllValidateError("UserAdminListInteractor.List", "validate income data failed", err)
	}

	res, err := ual.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, domerrs.NewBllError("UserAdminListInteractor.List", fmt.Sprintf("list User data with limit [%v] and offset [%v] failed", limit, offset), err)
	}

	return res, nil
}

func (ual *UserAdminListInteractor) validate(limit, offset int) error {
	// correct limit
	if limit <= 0 {
		return errs.NewInvalidArgumentError("limit", "must be greater than zero")
	}
	// max limit
	if limit > ual.maxListLimit {
		return errs.NewInvalidArgumentError("limit", fmt.Sprintf("must be less than or equal to max limit [%v]", ual.maxListLimit))
	}
	// offset
	if offset < 0 {
		return errs.NewInvalidArgumentError("offset", "must be greater or equal than zero")
	}

	return nil
}
