package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminGetUseCase defines the application business logic boundary for handling administrative user retrieval operations.
type UserAdminGetUseCase interface {
	// Get retrieves a specific user entity blueprint matching the provided unique identity sequence.
	Get(ctx context.Context, ID string) (*domain.User, error)
}

// UserAdminGetInteractor implements the UserAdminGetUseCase interface, orchestrating criteria evaluation
// and data storage lookup queries to deliver user identity structures.
type UserAdminGetInteractor struct {
	userRepo domain.UserAdminRepository // DAL administrative repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ UserAdminGetUseCase = (*UserAdminGetInteractor)(nil)

// NewUserAdminGetUseCase acts as a factory constructor mounting required user administration repository dependencies.
func NewUserAdminGetUseCase(userRepo domain.UserAdminRepository) *UserAdminGetInteractor {
	return &UserAdminGetInteractor{
		userRepo: userRepo,
	}
}

// Get validates inbound identification bounds and delegates execution to the persistence layer to fetch user data aggregate fields.
func (uag *UserAdminGetInteractor) Get(ctx context.Context, ID string) (*domain.User, error) {
	if err := uag.validate(ID); err != nil {
		return nil, errs.NewBllValidateError("UserAdminGetInteractor.Get", "validate income data failed", err)
	}

	res, err := uag.userRepo.Find(ctx, ID)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("UserAdminGetInteractor.Get", "User", ID, err)
		}

		return nil, errs.NewBllError("UserAdminGetInteractor.Get", fmt.Sprintf("find User model id [%s] failed", ID), err)
	}

	return res, nil
}

// validate executes initial syntax analysis over inbound query parameters before hitting storage boundaries.
func (uag *UserAdminGetInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
