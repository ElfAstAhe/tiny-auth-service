package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminGetNameUseCase defines the application business logic boundary for handling administrative user retrieval operations by name identifiers.
type UserAdminGetNameUseCase interface {
	// Get retrieves a specific user entity blueprint matching the provided unique string name key.
	Get(ctx context.Context, name string) (*domain.User, error)
}

// UserAdminGetNameInteractor implements the UserAdminGetNameUseCase interface, orchestrating criteria evaluation
// and data storage lookup queries to deliver user identity structures matching string keys.
type UserAdminGetNameInteractor struct {
	userRepo domain.UserAdminRepository // DAL administrative repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ UserAdminGetNameUseCase = (*UserAdminGetNameInteractor)(nil)

// NewUserAdminGetNameUseCase acts as a factory constructor mounting required user administration repository dependencies.
func NewUserAdminGetNameUseCase(userRepo domain.UserAdminRepository) *UserAdminGetNameInteractor {
	return &UserAdminGetNameInteractor{
		userRepo: userRepo,
	}
}

// Get validates inbound name criteria bounds and delegates execution to the persistence layer to fetch user data aggregate fields.
func (uag *UserAdminGetNameInteractor) Get(ctx context.Context, name string) (*domain.User, error) {
	if err := uag.validate(name); err != nil {
		return nil, errs.NewBllValidateError("UserAdminGetNameInteractor.Get", "validate income data failed", err)
	}

	res, err := uag.userRepo.FindByName(ctx, name)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("UserAdminGetNameInteractor.Get", "User", name, err)
		}

		return nil, errs.NewBllError("UserAdminGetNameInteractor.Get", fmt.Sprintf("find User model name [%s] failed", name), err)
	}

	return res, nil
}

// validate executes initial presence analysis over inbound name query criteria before hitting storage.
func (uag *UserAdminGetNameInteractor) validate(name string) error {
	if strings.TrimSpace(name) == "" {
		return errs.NewInvalidArgumentError("name", "name is empty")
	}

	return nil
}
