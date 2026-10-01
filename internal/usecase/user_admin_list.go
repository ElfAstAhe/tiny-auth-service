package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminListUseCase defines the application business logic boundary for handling administrative user collection queries.
type UserAdminListUseCase interface {
	// List retrieves a paginated collection of user entities matching boundary criteria metrics.
	List(ctx context.Context, limit, offset int) ([]*domain.User, error)
}

// UserAdminListInteractor implements the UserAdminListUseCase interface, orchestrating data pagination criteria constraints,
// limit validation, and repository collection lookups.
type UserAdminListInteractor struct {
	userRepo     domain.UserAdminRepository // DAL administrative repository handle for user state persistence
	maxListLimit int                        // Upper safety boundary limit enforced to prevent memory allocation spikes
}

// Compile-time interface compliance verification
var _ UserAdminListUseCase = (*UserAdminListInteractor)(nil)

// NewUserAdminListUseCase acts as a factory constructor mounting required user administration repository and safety thresholds.
func NewUserAdminListUseCase(userRepo domain.UserAdminRepository, maxListLimit int) *UserAdminListInteractor {
	res := &UserAdminListInteractor{
		userRepo:     userRepo,
		maxListLimit: maxListLimit,
	}
	if res.maxListLimit < 0 {
		res.maxListLimit = DefaultMaxLimit
	}

	return res
}

// List validates input pagination parameters and queries the data adapter layer to fetch sequential user array elements.
func (ual *UserAdminListInteractor) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	if err := ual.validate(limit, offset); err != nil {
		return nil, errs.NewBllValidateError("UserAdminListInteractor.List", "validate income data failed", err)
	}

	res, err := ual.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, errs.NewBllError("UserAdminListInteractor.List", fmt.Sprintf("list User data with limit [%v] and offset [%v] failed", limit, offset), err)
	}

	return res, nil
}

// validate executes semantic range and integrity checks on input pagination criteria before hitting storage.
func (ual *UserAdminListInteractor) validate(limit, offset int) error {
	// correct limit
	if limit <= 0 {
		return errs.NewInvalidArgumentError("limit", "must be greater than zero")
	}
	// max limit
	if limit > ual.maxListLimit {
		return errs.NewInvalidArgumentError("limit", fmt.Sprintf("must be less than or equal to max limit [%v]", ual.maxListLimit))
	}
	// offset
	if offset < 0 {
		return errs.NewInvalidArgumentError("offset", "must be greater or equal than zero")
	}

	return nil
}
