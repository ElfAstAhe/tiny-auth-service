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

type RoleAdminDeleteUseCase interface {
	Delete(context.Context, string) error
}

type RoleAdminDeleteInteractor struct {
	tm       usecase.TransactionManager
	roleRepo domain.RoleAdminRepository
}

var _ RoleAdminDeleteUseCase = (*RoleAdminDeleteInteractor)(nil)

func NewRoleAdminDeleteUseCase(tm usecase.TransactionManager, roleRepo domain.RoleAdminRepository) *RoleAdminDeleteInteractor {
	return &RoleAdminDeleteInteractor{
		tm:       tm,
		roleRepo: roleRepo,
	}
}

func (rad *RoleAdminDeleteInteractor) Delete(ctx context.Context, ID string) error {
	if err := rad.validate(ID); err != nil {
		return domerrs.NewBllValidateError("RoleAdminDeleteInteractor.Delete", "validate income data failed", err)
	}

	err := rad.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		return rad.roleRepo.Delete(ctx, ID)
	})
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return domerrs.NewBllNotFoundError("RoleAdminDeleteInteractor.Delete", "Role", ID, err)
		}

		return domerrs.NewBllError("RoleAdminDeleteInteractor.Delete", fmt.Sprintf("delete role model id [%v] failed", ID), err)
	}

	return nil
}

func (rad *RoleAdminDeleteInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
