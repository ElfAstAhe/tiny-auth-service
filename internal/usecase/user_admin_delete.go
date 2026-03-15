package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	usecase "github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type UserAdminDeleteUseCase interface {
	Delete(ctx context.Context, ID string) error
}

type UserAdminDeleteInteractor struct {
	tm       usecase.TransactionManager
	userRepo domain.UserAdminRepository
}

func NewUserAdminDeleteUseCase(tm usecase.TransactionManager, userRepo domain.UserAdminRepository) *UserAdminDeleteInteractor {
	return &UserAdminDeleteInteractor{
		tm:       tm,
		userRepo: userRepo,
	}
}

func (uad *UserAdminDeleteInteractor) Delete(ctx context.Context, ID string) error {
	if err := uad.validate(ID); err != nil {
		return domerrs.NewBllValidateError("UserAdminDeleteInteractor.Delete", "validate income data failed", err)
	}

	err := uad.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		return uad.userRepo.Delete(ctx, ID)
	})
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return domerrs.NewBllNotFoundError("UserAdminDeleteInteractor.Delete", "User", ID, err)
		}

		return domerrs.NewBllError("UserAdminDeleteInteractor.Delete", fmt.Sprintf("delete User model id [%s] failed", ID), err)
	}

	return nil
}

func (uad *UserAdminDeleteInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
