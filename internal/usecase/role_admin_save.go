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

type RoleAdminSaveUseCase interface {
	Save(ctx context.Context, model *domain.Role) (*domain.Role, error)
}

type RoleAdminSaveInteractor struct {
	tm       usecase.TransactionManager
	roleRepo domain.RoleAdminRepository
}

var _ RoleAdminSaveUseCase = (*RoleAdminSaveInteractor)(nil)

func NewRoleAdminSaveUseCase(tm usecase.TransactionManager, roleRepo domain.RoleAdminRepository) *RoleAdminSaveInteractor {
	return &RoleAdminSaveInteractor{
		tm:       tm,
		roleRepo: roleRepo,
	}
}

func (ras *RoleAdminSaveInteractor) Save(ctx context.Context, model *domain.Role) (*domain.Role, error) {
	var res *domain.Role
	err := ras.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		var txErr error
		if !model.IsExists() {
			res, txErr = ras.roleRepo.Create(ctx, model)
		} else {
			res, txErr = ras.roleRepo.Change(ctx, model)
		}

		return txErr
	})
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return nil, domerrs.NewBllNotFoundError("RoleAdminSaveInteractor.Save", "Role", model.ID, err)
		}
		if errors.As(err, new(*errs.DalAlreadyExistsError)) {
			return nil, domerrs.NewBllUniqueError("RoleAdminSaveInteractor.Save", "Role", model.ID, err)
		}

		return nil, domerrs.NewBllError("RoleAdminSaveInteractor.Save", fmt.Sprintf("save Role model id [%v] failed", model.ID), err)
	}

	return res, nil
}
