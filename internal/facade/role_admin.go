package facade

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

// RoleAdminFacade defines the administrative orchestration contract managing role records mapping and RBAC safety checkpoints.
type RoleAdminFacade interface {
	Get(ctx context.Context, ID string) (*dto.RoleDTO, error)
	GetByName(ctx context.Context, name string) (*dto.RoleDTO, error)
	List(ctx context.Context, limit, offset int) ([]*dto.RoleDTO, error)
	Create(ctx context.Context, role *dto.RoleDTO) (*dto.RoleDTO, error)
	Change(ctx context.Context, ID string, role *dto.RoleDTO) (*dto.RoleDTO, error)
	Delete(ctx context.Context, ID string) error
}

// RoleAdminFacadeImpl structures administrative boundary routers, conducting data transformations and permission assertions.
type RoleAdminFacadeImpl struct {
	authHelper   auth.Helper                     // Framework core token management tool verifying structural context identity claims
	getUC        usecase.RoleAdminGetUseCase     // Downstream application usecase handling administrative specific identifier lookup
	getByNameUC  usecase.RoleAdminGetNameUseCase // Downstream application usecase managing criteria name string lookup routines
	listUC       usecase.RoleAdminListUseCase    // Downstream application usecase managing array paginated aggregate lookup queries
	saveUC       usecase.RoleAdminSaveUseCase    // Downstream application usecase coordinating persistence operations across creation and update workflows
	deleteUC     usecase.RoleAdminDeleteUseCase  // Downstream application usecase executing single identity entity eviction bounds
	maxListLimit int                             // Safety threshold boundary limit preventing memory exhaustion during extensive query operations
}

// Compile-time interface compliance verification
var _ RoleAdminFacade = (*RoleAdminFacadeImpl)(nil)

// NewRoleAdminFacade acts as a factory constructor embedding granular administrative role usecase components and framework helpers.
func NewRoleAdminFacade(
	authHelper auth.Helper,
	getUC usecase.RoleAdminGetUseCase,
	getByNameUC usecase.RoleAdminGetNameUseCase,
	listUC usecase.RoleAdminListUseCase,
	saveUC usecase.RoleAdminSaveUseCase,
	deleteUC usecase.RoleAdminDeleteUseCase,
	maxListLimit int,
) *RoleAdminFacadeImpl {
	return &RoleAdminFacadeImpl{
		authHelper:   authHelper,
		getUC:        getUC,
		getByNameUC:  getByNameUC,
		listUC:       listUC,
		saveUC:       saveUC,
		deleteUC:     deleteUC,
		maxListLimit: maxListLimit,
	}
}

// Get executes security credential checks and translates internal role domain model data payloads back into external transfer objects by ID keys.
func (raf *RoleAdminFacadeImpl) Get(ctx context.Context, ID string) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.Get", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.Get", "user is not an admin", err)
	}

	if strings.TrimSpace(ID) == "" {
		return nil, errs.NewInvalidArgumentError("ID", "id is required")
	}

	model, err := raf.getUC.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

// GetByName evaluates identity roles before requesting the target entity array matching specific name query sequences.
func (raf *RoleAdminFacadeImpl) GetByName(ctx context.Context, name string) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.GetByName", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.GetByName", "user is not an admin", err)
	}

	if strings.TrimSpace(name) == "" {
		return nil, errs.NewInvalidArgumentError("name", "name is required")
	}

	model, err := raf.getByNameUC.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

// List reviews paginated range boundaries and returns a transformed decoupled collection array target.
func (raf *RoleAdminFacadeImpl) List(ctx context.Context, limit, offset int) ([]*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.List", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.List", "user is not an admin", err)
	}

	if err := raf.validateList(limit, offset); err != nil {
		return nil, err
	}

	models, err := raf.listUC.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return mapper.MapRolesModelToDTO(models), nil
}

// validateList asserts structural integrity of input pagination indices prior to executing resource operations.
func (raf *RoleAdminFacadeImpl) validateList(limit, offset int) error {
	if limit <= 0 {
		return errs.NewInvalidArgumentError("limit", "must be greater than 0")
	}
	if offset < 0 {
		return errs.NewInvalidArgumentError("offset", "must be greater or equal than 0")
	}
	if limit > raf.maxListLimit {
		return errs.NewInvalidArgumentError("limit", fmt.Sprintf("must be less or equal than %v", raf.maxListLimit))
	}

	return nil
}

// Create enforces admin restrictions, maps external data payloads onto domain schemas, and delegates atomic insertion routines.
func (raf *RoleAdminFacadeImpl) Create(ctx context.Context, role *dto.RoleDTO) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.Create", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.Create", "user is not an admin", err)
	}

	if role == nil {
		return nil, errs.NewInvalidArgumentError("role", "is required")
	}

	model := mapper.MapRoleDTOToModel(role)
	model.ID = ""

	model, err = raf.saveUC.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

// Change ensures strict authorization mappings, updates domain models with context identifier keys, and commits changes downstream.
func (raf *RoleAdminFacadeImpl) Change(ctx context.Context, ID string, role *dto.RoleDTO) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.Change", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("RoleAdminFacadeImpl.Change", "user is not an admin", err)
	}

	if role == nil {
		return nil, errs.NewInvalidArgumentError("role", "is required")
	}

	model := mapper.MapRoleDTOToModel(role)
	model.ID = ID

	model, err = raf.saveUC.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

// Delete acts as a destructive gateway validation checkpoint before passing eviction criteria commands down the line.
func (raf *RoleAdminFacadeImpl) Delete(ctx context.Context, ID string) error {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return errs.NewBllForbiddenError("RoleAdminFacadeImpl.Delete", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return errs.NewBllForbiddenError("RoleAdminFacadeImpl.Delete", "user is not an admin", err)
	}

	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is required")
	}

	return raf.deleteUC.Delete(ctx, ID)
}
