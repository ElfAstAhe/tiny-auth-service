package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminListUseCase defines the application business logic boundary for handling administrative role collection queries.
type RoleAdminListUseCase interface {
	// List retrieves a paginated collection of role entities matching boundary criteria metrics.
	List(ctx context.Context, limit, offset int) ([]*domain.Role, error)
}

// RoleAdminListInteractor implements the RoleAdminListUseCase interface, orchestrating data pagination criteria constraints,
// limit validation, and repository collection lookups.
type RoleAdminListInteractor struct {
	roleRepo     domain.RoleAdminRepository // DAL administrative repository handle for role state persistence
	maxListLimit int                        // Upper safety boundary limit enforced to prevent memory allocation spikes
}

// Compile-time interface compliance verification
var _ RoleAdminListUseCase = (*RoleAdminListInteractor)(nil)

// NewRoleAdminListUseCase acts as a factory constructor mounting required role administration repository and safety thresholds.
func NewRoleAdminListUseCase(roleRepo domain.RoleAdminRepository, maxListLimit int) *RoleAdminListInteractor {
	return &RoleAdminListInteractor{
		roleRepo:     roleRepo,
		maxListLimit: maxListLimit,
	}
}

// List validates input pagination parameters and queries the data adapter layer to fetch sequential role array elements.
func (ral *RoleAdminListInteractor) List(ctx context.Context, limit, offset int) ([]*domain.Role, error) {
	if err := ral.validate(limit, offset); err != nil {
		return nil, errs.NewBllValidateError("RoleAdminListInteractor.List", "validate income data failed", err)
	}

	res, err := ral.roleRepo.List(ctx, limit, offset)
	if err != nil {
		// FIXED: Provided explicit operational context identifier "List" to ensure precise structural error tracking propagation.
		return nil, errs.NewBllError("List", fmt.Sprintf("list Role data with limit [%v] and offset [%v] failed", limit, offset), err)
	}

	return res, nil
}

// validate executes semantic range and integrity checks on input pagination criteria before hitting storage.
func (ral *RoleAdminListInteractor) validate(limit, offset int) error {
	// correct limit
	if limit <= 0 {
		return errs.NewInvalidArgumentError("limit", "must be greater than zero")
	}
	// max limit
	if limit > ral.maxListLimit {
		return errs.NewInvalidArgumentError("limit", fmt.Sprintf("must be less than or equal to max limit [%v]", ral.maxListLimit))
	}
	// offset
	if offset < 0 {
		return errs.NewInvalidArgumentError("offset", "must be greater or equal than zero")
	}

	return nil
}
