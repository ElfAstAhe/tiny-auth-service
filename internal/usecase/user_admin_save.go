package usecase

import (
	"context"
	"errors"
	"fmt"

	usecase "github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type UserAdminSaveUseCase interface {
	Save(ctx context.Context, model *domain.User) (*domain.User, error)
}

type UserAdminSaveInteractor struct {
	tm       usecase.TransactionManager
	userRepo domain.UserAdminRepository
}

func NewUserAdminSaveUseCase(tm usecase.TransactionManager, userRepo domain.UserAdminRepository) *UserAdminSaveInteractor {
	return &UserAdminSaveInteractor{
		tm:       tm,
		userRepo: userRepo,
	}
}

func (uas *UserAdminSaveInteractor) Save(ctx context.Context, model *domain.User) (*domain.User, error) {
	var res *domain.User
	err := uas.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		var txErr error
		if !model.IsExists() {
			res, txErr = uas.userRepo.Create(ctx, model)
		} else {
			res, txErr = uas.userRepo.Change(ctx, model)
		}

		return txErr
	})
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return nil, domerrs.NewBllNotFoundError("UserAdminSaveInteractor.Save", "User", model.ID, err)
		}
		if errors.As(err, new(*errs.DalAlreadyExistsError)) {
			return nil, domerrs.NewBllUniqueError("UserAdminSaveInteractor.Save", "User", model.ID, err)
		}

		return nil, domerrs.NewBllError("UserAdminSaveInteractor.Save", fmt.Sprintf("save User model id [%v] failed", model.ID), err)
	}

	return res, err
}
