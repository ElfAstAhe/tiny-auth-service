package facade

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

type RoleAdminFacade interface {
	Get(ctx context.Context, ID string) (*dto.RoleDto, error)
	GetByName(ctx context.Context, name string) (*dto.RoleDto, error)
	List(ctx context.Context, limit, offset int) ([]*dto.RoleDto, error)
	Create(ctx context.Context, role *dto.RoleDto) (*dto.RoleDto, error)
	Change(ctx context.Context, ID string, role *dto.RoleDto) (*dto.RoleDto, error)
	Delete(ctx context.Context, ID string) error
}

type RoleAdminFacadeImpl struct {
	getUC        usecase.RoleAdminGetUseCase
	getByNameUC  usecase.RoleAdminGetUseCase
	listUC       usecase.RoleAdminListUseCase
	saveUC       usecase.RoleAdminSaveUseCase
	deleteUC     usecase.RoleAdminDeleteUseCase
	maxListLimit int
}

func NewRoleAdminFacade(
	getUC usecase.RoleAdminGetUseCase,
	getByNameUC usecase.RoleAdminGetUseCase,
	listUC usecase.RoleAdminListUseCase,
	saveUC usecase.RoleAdminSaveUseCase,
	deleteUC usecase.RoleAdminDeleteUseCase,
	maxListLimit int,
) *RoleAdminFacadeImpl {
	return &RoleAdminFacadeImpl{
		getUC:        getUC,
		getByNameUC:  getByNameUC,
		listUC:       listUC,
		saveUC:       saveUC,
		deleteUC:     deleteUC,
		maxListLimit: maxListLimit,
	}
}

func (raf *RoleAdminFacadeImpl) Get(ctx context.Context, ID string) (*dto.RoleDto, error) {
	if strings.TrimSpace(ID) == "" {
		return nil, errs.NewInvalidArgumentError("ID", "id is required")
	}

	model, err := raf.getUC.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

func (raf *RoleAdminFacadeImpl) GetByName(ctx context.Context, name string) (*dto.RoleDto, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errs.NewInvalidArgumentError("name", "name is required")
	}

	model, err := raf.getByNameUC.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

func (raf *RoleAdminFacadeImpl) List(ctx context.Context, limit, offset int) ([]*dto.RoleDto, error) {
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

func (raf *RoleAdminFacadeImpl) Create(ctx context.Context, role *dto.RoleDto) (*dto.RoleDto, error) {
	if role == nil {
		return nil, errs.NewInvalidArgumentError("role", "is required")
	}

	model := mapper.MapRoleDTOToModel(role)
	model.ID = ""

	var err error
	model, err = raf.saveUC.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

func (raf *RoleAdminFacadeImpl) Change(ctx context.Context, ID string, role *dto.RoleDto) (*dto.RoleDto, error) {
	if role == nil {
		return nil, errs.NewInvalidArgumentError("role", "is required")
	}

	model := mapper.MapRoleDTOToModel(role)
	model.ID = ID

	var err error
	model, err = raf.saveUC.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.MapRoleModelToDTO(model), nil
}

func (raf *RoleAdminFacadeImpl) Delete(ctx context.Context, ID string) error {
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is required")
	}

	return raf.deleteUC.Delete(ctx, ID)
}
