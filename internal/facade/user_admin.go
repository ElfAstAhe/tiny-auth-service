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

// UserAdminFacade defines the administrative orchestration contract managing user records mapping and RBAC safety checkpoints.
type UserAdminFacade interface {
	Get(ctx context.Context, ID string) (*dto.UserDTO, error)
	GetByName(ctx context.Context, name string) (*dto.UserDTO, error)
	List(ctx context.Context, limit, offset int) ([]*dto.UserDTO, error)
	Create(ctx context.Context, user *dto.UserDTO) (*dto.UserDTO, error)
	Change(ctx context.Context, ID string, user *dto.UserDTO) (*dto.UserDTO, error)
	Delete(ctx context.Context, ID string) error
}

// UserAdminFacadeImpl structures administrative boundary routers, conducting data transformations and permission assertions.
type UserAdminFacadeImpl struct {
	authHelper   auth.Helper                     // Framework core token management tool verifying structural context identity claims
	getUC        usecase.UserAdminGetUseCase     // Downstream application usecase handling administrative specific identifier lookup
	getByNameUC  usecase.UserAdminGetNameUseCase // Downstream application usecase managing criteria name string lookup routines
	listUC       usecase.UserAdminListUseCase    // Downstream application usecase managing array paginated aggregate lookup queries
	saveUC       usecase.UserAdminSaveUseCase    // Downstream application usecase coordinating persistence operations across creation and update workflows
	deleteUC     usecase.UserAdminDeleteUseCase  // Downstream application usecase executing single identity entity eviction bounds
	maxListLimit int                             // Safety threshold boundary limit preventing memory exhaustion during extensive query operations
}

// Compile-time interface compliance verification
var _ UserAdminFacade = (*UserAdminFacadeImpl)(nil)

// NewUserAdminFacade acts as a factory constructor embedding granular administrative usecase components and framework helpers.
func NewUserAdminFacade(
	authHelper auth.Helper,
	getUC usecase.UserAdminGetUseCase,
	getByNameUC usecase.UserAdminGetNameUseCase,
	listUC usecase.UserAdminListUseCase,
	saveUC usecase.UserAdminSaveUseCase,
	deleteUC usecase.UserAdminDeleteUseCase,
	maxListLimit int,
) *UserAdminFacadeImpl {
	return &UserAdminFacadeImpl{
		authHelper:   authHelper,
		getUC:        getUC,
		getByNameUC:  getByNameUC,
		listUC:       listUC,
		saveUC:       saveUC,
		deleteUC:     deleteUC,
		maxListLimit: maxListLimit,
	}
}

// Get executes security credential checks and translates internal user domain model data payloads back into external transfer objects by ID keys.
func (uaf *UserAdminFacadeImpl) Get(ctx context.Context, ID string) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.Get", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.Get", "user is not an admin", err)
	}

	if strings.TrimSpace(ID) == "" {
		return nil, errs.NewInvalidArgumentError("ID", "id is required")
	}

	model, err := uaf.getUC.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToDTO(model), nil
}

// GetByName evaluates identity roles before requesting the target entity array matching specific name query sequences.
func (uaf *UserAdminFacadeImpl) GetByName(ctx context.Context, name string) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.GetByName", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.GetByName", "user is not an admin", err)
	}

	if strings.TrimSpace(name) == "" {
		return nil, errs.NewInvalidArgumentError("name", "name is required")
	}

	model, err := uaf.getByNameUC.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToDTO(model), nil
}

// List reviews paginated range boundaries and returns a transformed decoupled collection array target.
func (uaf *UserAdminFacadeImpl) List(ctx context.Context, limit, offset int) ([]*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.List", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.List", "user is not an admin", err)
	}

	if err := uaf.validateList(limit, offset); err != nil {
		return nil, err
	}

	models, err := uaf.listUC.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelsToDTO(models), nil
}

// validateList asserts structural integrity of input pagination indices prior to executing resource operations.
func (uaf *UserAdminFacadeImpl) validateList(limit, offset int) error {
	if limit <= 0 {
		return errs.NewInvalidArgumentError("limit", "must be greater than 0")
	}
	if offset < 0 {
		return errs.NewInvalidArgumentError("offset", "must be greater or equal than 0")
	}
	if limit > uaf.maxListLimit {
		return errs.NewInvalidArgumentError("limit", fmt.Sprintf("must be less or equal than %v", uaf.maxListLimit))
	}

	return nil
}

// Create enforces admin restrictions, maps external data payloads onto domain schemas, and delegates atomic insertion routines.
func (uaf *UserAdminFacadeImpl) Create(ctx context.Context, user *dto.UserDTO) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.Create", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.Create", "user is not an admin", err)
	}

	if user == nil {
		return nil, errs.NewInvalidArgumentError("user", "is required")
	}

	model := mapper.MapUserDTOToModel(user)
	model.ID = ""

	model, err = uaf.saveUC.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToDTO(model), nil
}

// Change ensures strict authorization mappings, updates domain models with context identifier keys, and commits changes downstream.
func (uaf *UserAdminFacadeImpl) Change(ctx context.Context, ID string, user *dto.UserDTO) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.Change", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, errs.NewBllForbiddenError("UserAdminFacadeImpl.Change", "user is not an admin", err)
	}

	if user == nil {
		return nil, errs.NewInvalidArgumentError("user", "is required")
	}

	model := mapper.MapUserDTOToModel(user)
	model.ID = ID

	model, err = uaf.saveUC.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToDTO(model), nil
}

// Delete acts as a destructive gateway validation checkpoint before passing eviction criteria commands down the line.
func (uaf *UserAdminFacadeImpl) Delete(ctx context.Context, ID string) error {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return errs.NewBllForbiddenError("UserAdminFacadeImpl.Delete", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return errs.NewBllForbiddenError("UserAdminFacadeImpl.Delete", "user is not an admin", err)
	}
	// validate income
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is required")
	}
	// logic
	return uaf.deleteUC.Delete(ctx, ID)
}
