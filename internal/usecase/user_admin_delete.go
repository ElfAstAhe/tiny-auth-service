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

// UserAdminDeleteUseCase defines the application business logic boundary for handling administrative user eviction scenarios.
type UserAdminDeleteUseCase interface {
	// Delete removes a user identity structure matching the specified identification sequence.
	Delete(ctx context.Context, ID string) error
}

// UserAdminDeleteInteractor implements the UserAdminDeleteUseCase interface, orchestrating data validation
// and entity eviction queries inside a managed unit of work database context.
type UserAdminDeleteInteractor struct {
	uw       libdom.UnitOfWork          // BLL-level unit of work boundary abstraction manager
	userRepo domain.UserAdminRepository // DAL administrative repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ UserAdminDeleteUseCase = (*UserAdminDeleteInteractor)(nil)

// NewUserAdminDeleteUseCase acts as a factory constructor mounting required user administration repository and transaction dependencies.
func NewUserAdminDeleteUseCase(
	uw libdom.UnitOfWork,
	userRepo domain.UserAdminRepository,
) *UserAdminDeleteInteractor {
	return &UserAdminDeleteInteractor{
		uw:       uw,
		userRepo: userRepo,
	}
}

// Delete validates criteria input bounds and delegates execution to the storage layer within an isolated transactional wrapper.
func (uad *UserAdminDeleteInteractor) Delete(ctx context.Context, ID string) error {
	if err := uad.validate(ID); err != nil {
		return errs.NewBllValidateError("UserAdminDeleteInteractor.Delete", "validate income data failed", err)
	}

	// FIXED: Utilizing unique 'txCtx' sequence parameter inside the closure callback to guarantee strict ACID compliance
	err := uad.uw.Execute(ctx, func(txCtx context.Context) error {
		return uad.userRepo.Delete(txCtx, ID)
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return errs.NewBllNotFoundError("UserAdminDeleteInteractor.Delete", "User", ID, err)
		}

		return errs.NewBllError("UserAdminDeleteInteractor.Delete", fmt.Sprintf("delete User model id [%s] failed", ID), err)
	}

	return nil
}

// validate executes initial syntax and presence checks over inbound identification payloads.
func (uad *UserAdminDeleteInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
