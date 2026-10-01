package usecase

import (
	"context"
	"errors"
	"fmt"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminSaveUseCase defines the application business logic boundary for handling administrative role persistence and modification workflows.
type RoleAdminSaveUseCase interface {
	// Save creates a new role identity or modifies an existing one within an atomic transaction boundary.
	Save(ctx context.Context, model *domain.Role) (*domain.Role, error)
}

// RoleAdminSaveInteractor implements the RoleAdminSaveUseCase interface, orchestrating dynamic operational routing
// between creation and updates, and mapping persistence operations inside a managed unit of work database context.
type RoleAdminSaveInteractor struct {
	uw       libdom.UnitOfWork          // BLL-level unit of work boundary abstraction manager
	roleRepo domain.RoleAdminRepository // DAL administrative repository handle for role state persistence
}

// Compile-time interface compliance verification
var _ RoleAdminSaveUseCase = (*RoleAdminSaveInteractor)(nil)

// NewRoleAdminSaveUseCase acts as a factory constructor mounting required role administration repository and transaction dependencies.
func NewRoleAdminSaveUseCase(uw libdom.UnitOfWork, roleRepo domain.RoleAdminRepository) *RoleAdminSaveInteractor {
	return &RoleAdminSaveInteractor{
		uw:       uw,
		roleRepo: roleRepo,
	}
}

// Save routes the operation based on identity presence and commits structural status criteria inside a transactional closure.
func (ras *RoleAdminSaveInteractor) Save(ctx context.Context, model *domain.Role) (*domain.Role, error) {
	var res *domain.Role
	// сохраняем
	// FIXED: Utilizing unique 'txCtx' sequence parameter inside the closure callback to guarantee strict ACID compliance
	err := ras.uw.Execute(ctx, func(txCtx context.Context) error {
		var txErr error
		if !model.IsExists() {
			res, txErr = ras.roleRepo.Create(txCtx, model)
		} else {
			res, txErr = ras.roleRepo.Change(txCtx, model)
		}

		return txErr
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("RoleAdminSaveInteractor.Save", "Role", model.ID, err)
		}
		if _, ok := errors.AsType[*errs.DalAlreadyExistsError](err); ok {
			return nil, errs.NewBllUniqueError("RoleAdminSaveInteractor.Save", "Role", model.ID, err)
		}

		return nil, errs.NewBllError("RoleAdminSaveInteractor.Save", fmt.Sprintf("save Role model id [%v] failed", model.ID), err)
	}

	return res, nil
}
