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

type UserAdminFacade interface {
	Get(ctx context.Context, ID string) (*dto.UserDTO, error)
	GetByName(ctx context.Context, name string) (*dto.UserDTO, error)
	List(ctx context.Context, limit, offset int) ([]*dto.UserDTO, error)
	Create(ctx context.Context, user *dto.UserDTO) (*dto.UserDTO, error)
	Change(ctx context.Context, ID string, user *dto.UserDTO) (*dto.UserDTO, error)
	Delete(ctx context.Context, ID string) error
}

type UserAdminFacadeImpl struct {
	authHelper   auth.Helper
	getUC        usecase.UserAdminGetUseCase
	getByNameUC  usecase.UserAdminGetNameUseCase
	listUC       usecase.UserAdminListUseCase
	saveUC       usecase.UserAdminSaveUseCase
	deleteUC     usecase.UserAdminDeleteUseCase
	maxListLimit int
}

var _ UserAdminFacade = (*UserAdminFacadeImpl)(nil)

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

func (uaf *UserAdminFacadeImpl) Get(ctx context.Context, ID string) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Get", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Get", "user is not an admin", err)
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

func (uaf *UserAdminFacadeImpl) GetByName(ctx context.Context, name string) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.GetByName", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.GetByName", "user is not an admin", err)
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

func (uaf *UserAdminFacadeImpl) List(ctx context.Context, limit, offset int) ([]*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.List", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.List", "user is not an admin", err)
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

func (uaf *UserAdminFacadeImpl) validateList(limit, offset int) error {
	if !(limit > 0) {
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

func (uaf *UserAdminFacadeImpl) Create(ctx context.Context, user *dto.UserDTO) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Create", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Create", "user is not an admin", err)
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

func (uaf *UserAdminFacadeImpl) Change(ctx context.Context, ID string, user *dto.UserDTO) (*dto.UserDTO, error) {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Change", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return nil, domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Change", "user is not an admin", err)
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

func (uaf *UserAdminFacadeImpl) Delete(ctx context.Context, ID string) error {
	// subject
	subj, err := uaf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Delete", "retrieve subject", err)
	}
	// RBAC
	if !IsSubjectAdmin(subj) {
		return domerrs.NewBllForbiddenError("UserAdminFacadeImpl.Delete", "user is not an admin", err)
	}
	// validate income
	if strings.TrimSpace(ID) == "" {
		return errs.NewInvalidArgumentError("ID", "id is required")
	}
	// logic
	return uaf.deleteUC.Delete(ctx, ID)
}
