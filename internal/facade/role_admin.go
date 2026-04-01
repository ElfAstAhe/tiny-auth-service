package facade

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

type RoleAdminFacade interface {
	Get(ctx context.Context, ID string) (*dto.RoleDTO, error)
	GetByName(ctx context.Context, name string) (*dto.RoleDTO, error)
	List(ctx context.Context, limit, offset int) ([]*dto.RoleDTO, error)
	Create(ctx context.Context, role *dto.RoleDTO) (*dto.RoleDTO, error)
	Change(ctx context.Context, ID string, role *dto.RoleDTO) (*dto.RoleDTO, error)
	Delete(ctx context.Context, ID string) error
}

type RoleAdminFacadeImpl struct {
	authHelper   auth.Helper
	getUC        usecase.RoleAdminGetUseCase
	getByNameUC  usecase.RoleAdminGetUseCase
	listUC       usecase.RoleAdminListUseCase
	saveUC       usecase.RoleAdminSaveUseCase
	deleteUC     usecase.RoleAdminDeleteUseCase
	maxListLimit int
}

var _ RoleAdminFacade = (*RoleAdminFacadeImpl)(nil)

func NewRoleAdminFacade(
	authHelper auth.Helper,
	getUC usecase.RoleAdminGetUseCase,
	getByNameUC usecase.RoleAdminGetUseCase,
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

func (raf *RoleAdminFacadeImpl) Get(ctx context.Context, ID string) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Get", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Get", "user is not an admin", err)
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

func (raf *RoleAdminFacadeImpl) GetByName(ctx context.Context, name string) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.GetByName", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.GetByName", "user is not an admin", err)
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

func (raf *RoleAdminFacadeImpl) List(ctx context.Context, limit, offset int) ([]*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.List", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.List", "user is not an admin", err)
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

func (raf *RoleAdminFacadeImpl) validateList(limit, offset int) error {
	if !(limit > 0) {
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

func (raf *RoleAdminFacadeImpl) Create(ctx context.Context, role *dto.RoleDTO) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Create", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Create", "user is not an admin", err)
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

func (raf *RoleAdminFacadeImpl) Change(ctx context.Context, ID string, role *dto.RoleDTO) (*dto.RoleDTO, error) {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Change", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Change", "user is not an admin", err)
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

func (raf *RoleAdminFacadeImpl) Delete(ctx context.Context, ID string) error {
	// subject
	subj, err := raf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Delete", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return domerrs.NewBllForbiddenError("RoleAdminFacadeImpl.Delete", "user is not an admin", err)
	}

	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is required")
	}

	return raf.deleteUC.Delete(ctx, ID)
}
