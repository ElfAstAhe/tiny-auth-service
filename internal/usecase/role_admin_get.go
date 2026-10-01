package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminGetUseCase defines the application business logic boundary for handling administrative role retrieval operations.
type RoleAdminGetUseCase interface {
	// Get retrieves a specific role entity blueprint matching the provided unique identity sequence.
	Get(ctx context.Context, ID string) (*domain.Role, error)
}

// RoleAdminGetInteractor implements the RoleAdminGetUseCase interface, orchestrating criteria evaluation
// and data storage lookup queries to deliver role identity structures.
type RoleAdminGetInteractor struct {
	roleRepo domain.RoleAdminRepository // DAL administrative repository handle for role state persistence
}

// Compile-time interface compliance verification
var _ RoleAdminGetUseCase = (*RoleAdminGetInteractor)(nil)

// NewRoleAdminGetUseCase acts as a factory constructor mounting required role administration repository dependencies.
func NewRoleAdminGetUseCase(roleRepo domain.RoleAdminRepository) *RoleAdminGetInteractor {
	return &RoleAdminGetInteractor{
		roleRepo: roleRepo,
	}
}

// Get validates inbound identification bounds and delegates execution to the persistence layer to fetch role aggregate fields.
func (rag *RoleAdminGetInteractor) Get(ctx context.Context, ID string) (*domain.Role, error) {
	if err := rag.validate(ID); err != nil {
		return nil, errs.NewBllValidateError("RoleAdminGetInteractor.Get", "validate income data failed", err)
	}

	res, err := rag.roleRepo.Find(ctx, ID)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("RoleAdminGetInteractor.Get", "Role", ID, err)
		}

		return nil, errs.NewBllError("RoleAdminGetInteractor.Get", fmt.Sprintf("find Role model id [%s] failed", ID), err)
	}

	return res, nil
}

// validate executes initial syntax analysis over inbound query parameters before hitting storage boundaries.
func (rag *RoleAdminGetInteractor) validate(ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is empty")
	}

	return nil
}
