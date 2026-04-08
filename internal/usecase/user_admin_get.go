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

type UserAdminGetUseCase interface {
	Get(ctx context.Context, ID string) (*domain.User, error)
}

type UserAdminGetInteractor struct {
	userRepo domain.UserAdminRepository
}

var _ UserAdminGetUseCase = (*UserAdminGetInteractor)(nil)

func NewUserAdminGetUseCase(userRepo domain.UserAdminRepository) *UserAdminGetInteractor {
	return &UserAdminGetInteractor{
		userRepo: userRepo,
	}
}

func (uag *UserAdminGetInteractor) Get(ctx context.Context, ID string) (*domain.User, error) {
	if err := uag.validate(ID); err != nil {
		return nil, domerrs.NewBllValidateError("UserAdminGetInteractor.Get", "validate income data failed", err)
	}

	res, err := uag.userRepo.Find(ctx, ID)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, domerrs.NewBllNotFoundError("UserAdminGetInteractor.Get", "User", ID, err)
		}

		return nil, domerrs.NewBllError("UserAdminGetInteractor.Get", fmt.Sprintf("find User model id [%s] failed", ID), err)
	}

	return res, nil
}

func (uag *UserAdminGetInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
