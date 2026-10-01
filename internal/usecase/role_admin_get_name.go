package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminGetNameUseCase defines the application business logic boundary for handling administrative role retrieval operations by name identifiers.
type RoleAdminGetNameUseCase interface {
	// Get retrieves a specific role entity blueprint matching the provided unique string name key.
	Get(ctx context.Context, name string) (*domain.Role, error)
}

// RoleAdminGetNameInteractor implements the RoleAdminGetNameUseCase interface, orchestrating criteria evaluation
// and data storage lookup queries to deliver role identity structures matching string keys.
type RoleAdminGetNameInteractor struct {
	roleRepo domain.RoleAdminRepository // DAL administrative repository handle for role state persistence
}

// Compile-time interface compliance verification
var _ RoleAdminGetNameUseCase = (*RoleAdminGetNameInteractor)(nil)

// NewRoleAdminGetNameUseCase acts as a factory constructor mounting required role administration repository dependencies.
func NewRoleAdminGetNameUseCase(roleRepo domain.RoleAdminRepository) *RoleAdminGetNameInteractor {
	return &RoleAdminGetNameInteractor{
		roleRepo: roleRepo,
	}
}

// Get validates inbound name criteria bounds and delegates execution to the persistence layer to fetch role aggregate fields.
func (agn *RoleAdminGetNameInteractor) Get(ctx context.Context, name string) (*domain.Role, error) {
	if err := agn.validate(name); err != nil {
		return nil, errs.NewBllValidateError("RoleAdminGetNameInteractor.Get", "validate income data failed", err)
	}

	res, err := agn.roleRepo.FindByName(ctx, name)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("RoleAdminGetNameInteractor.Get", "Role", name, err)
		}

		return nil, errs.NewBllError("RoleAdminGetNameInteractor.Get", fmt.Sprintf("find Role model name [%s] failed", name), err)
	}

	return res, nil
}

// validate executes initial presence analysis over inbound name query criteria before hitting storage.
func (agn *RoleAdminGetNameInteractor) validate(name string) error {
	if strings.TrimSpace(name) == "" {
		return errs.NewInvalidArgumentError("name", "name is empty")
	}

	return nil
}
