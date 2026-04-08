package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type UserAdminGetNameUseCase interface {
	Get(ctx context.Context, name string) (*domain.User, error)
}

type UserAdminGetNameInteractor struct {
	userRepo domain.UserAdminRepository
}

var _ UserAdminGetNameUseCase = (*UserAdminGetNameInteractor)(nil)

func NewUserAdminGetNameUseCase(userRepo domain.UserAdminRepository) *UserAdminGetNameInteractor {
	return &UserAdminGetNameInteractor{
		userRepo: userRepo,
	}
}

func (uag *UserAdminGetNameInteractor) Get(ctx context.Context, name string) (*domain.User, error) {
	if err := uag.validate(name); err != nil {
		return nil, domerrs.NewBllValidateError("UserAdminGetNameInteractor.Get", "validate income data failed", err)
	}

	res, err := uag.userRepo.FindByName(ctx, name)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, domerrs.NewBllNotFoundError("UserAdminGetNameInteractor.Get", "User", name, err)
		}

		return nil, domerrs.NewBllError("UserAdminGetNameInteractor.Get", fmt.Sprintf("find User model name [%s] failed", name), err)
	}

	return res, nil
}

func (uag *UserAdminGetNameInteractor) validate(name string) error {
	if strings.TrimSpace(name) == "" {
		return errs.NewInvalidArgumentError("name", "name is empty")
	}

	return nil
}
