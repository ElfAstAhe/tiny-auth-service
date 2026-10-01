package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminDeleteUseCase defines the application business logic boundary for handling administrative role eviction scenarios.
type RoleAdminDeleteUseCase interface {
	// Delete removes a role entity structure matching the specified identification sequence.
	Delete(context.Context, string) error
}

// RoleAdminDeleteInteractor implements the RoleAdminDeleteUseCase interface, orchestrating data validation
// and entity eviction queries inside a managed unit of work database context.
type RoleAdminDeleteInteractor struct {
	uw       libdom.UnitOfWork          // BLL-level unit of work boundary abstraction manager
	roleRepo domain.RoleAdminRepository // DAL administrative repository handle for role state persistence
}

// Compile-time interface compliance verification
var _ RoleAdminDeleteUseCase = (*RoleAdminDeleteInteractor)(nil)

// NewRoleAdminDeleteUseCase acts as a factory constructor mounting required role administration repository and transaction dependencies.
func NewRoleAdminDeleteUseCase(
	uw libdom.UnitOfWork,
	roleRepo domain.RoleAdminRepository,
) *RoleAdminDeleteInteractor {
	return &RoleAdminDeleteInteractor{
		uw:       uw,
		roleRepo: roleRepo,
	}
}

// Delete validates criteria input bounds and delegates execution to the storage layer within an isolated transactional wrapper.
func (rad *RoleAdminDeleteInteractor) Delete(ctx context.Context, ID string) error {
	if err := rad.validate(ID); err != nil {
		return errs.NewBllValidateError("RoleAdminDeleteInteractor.Delete", "validate income data failed", err)
	}

	// FIXED: Utilizing unique 'txCtx' sequence parameter inside the closure callback to guarantee strict ACID compliance
	err := rad.uw.Execute(ctx, func(txCtx context.Context) error {
		return rad.roleRepo.Delete(txCtx, ID)
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return errs.NewBllNotFoundError("RoleAdminDeleteInteractor.Delete", "Role", ID, err)
		}

		return errs.NewBllError("RoleAdminDeleteInteractor.Delete", fmt.Sprintf("delete role model id [%v] failed", ID), err)
	}

	return nil
}

// validate executes initial syntax and presence checks over inbound identification payloads.
func (rad *RoleAdminDeleteInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
