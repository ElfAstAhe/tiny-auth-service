package container

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/audit"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

//goland:noinspection DuplicatedCode
func (fc *FacadeContainer) providerAuthFacade() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	jwtHelperInst, err := container.GetInstance[*helper.JWTHelper](InstanceJWTHelper)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	authAuditClientInst, err := container.GetInstance[client.AuthAuditClient](InstanceAuthAuditClient)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	loginUCInst, err := container.GetInstance[usecase.LoginUseCase](InstanceLoginUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	loginSimpleUCInst, err := container.GetInstance[usecase.LoginSimpleUseCase](InstanceLoginSimpleUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}

	return audit.NewAuthFacade(
		authAuditClientInst,
		confInst.App.NodeName,
		facade.NewAuthFacade(
			jwtHelperInst,
			loginUCInst,
			loginSimpleUCInst,
		)), nil
}

//goland:noinspection DuplicatedCode
func (fc *FacadeContainer) providerRoleAdminFacade() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminGetUCInst, err := container.GetInstance[usecase.RoleAdminGetUseCase](InstanceRoleAdminGetUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminGetByNameUCInst, err := container.GetInstance[usecase.RoleAdminGetNameUseCase](InstanceRoleAdminGetByNameUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminListUCInst, err := container.GetInstance[usecase.RoleAdminListUseCase](InstanceRoleAdminListUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminSaveUCInst, err := container.GetInstance[usecase.RoleAdminSaveUseCase](InstanceRoleAdminSaveUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminDeleteUCInst, err := container.GetInstance[usecase.RoleAdminDeleteUseCase](InstanceRoleAdminDeleteUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}

	return facade.NewRoleAdminFacade(
		authHelperInst,
		roleAdminGetUCInst,
		roleAdminGetByNameUCInst,
		roleAdminListUCInst,
		roleAdminSaveUCInst,
		roleAdminDeleteUCInst,
		confInst.App.MaxListLimit,
	), nil
}

//goland:noinspection DuplicatedCode
func (fc *FacadeContainer) providerUserFacade() (any, error) {
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	registerUCInst, err := container.GetInstance[usecase.RegisterUseCase](InstanceRegisterUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	profileUCInst, err := container.GetInstance[usecase.ProfileUseCase](InstanceProfileUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	changePasswordUCInst, err := container.GetInstance[usecase.ChangePasswordUseCase](InstanceChangePasswordUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	changeKeysUCInst, err := container.GetInstance[usecase.ChangeKeysUseCase](InstanceChangeKeysUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}

	return facade.NewUserFacade(
		authHelperInst,
		registerUCInst,
		profileUCInst,
		changePasswordUCInst,
		changeKeysUCInst,
	), nil
}

//goland:noinspection DuplicatedCode
func (fc *FacadeContainer) providerUserAdminFacade() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminGetUCInst, err := container.GetInstance[usecase.UserAdminGetUseCase](InstanceUserAdminGetUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminGetByNameUCInst, err := container.GetInstance[usecase.UserAdminGetNameUseCase](InstanceUserAdminGetByNameUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminListUCInst, err := container.GetInstance[usecase.UserAdminListUseCase](InstanceUserAdminListUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminSaveUCInst, err := container.GetInstance[usecase.UserAdminSaveUseCase](InstanceUserAdminSaveUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminDeleteUCInst, err := container.GetInstance[usecase.UserAdminDeleteUseCase](InstanceUserAdminDeleteUC)
	if err != nil {
		return nil, errs.NewContainerError(fc.GetName(), "provider: retrieve instance failed", err)
	}

	return facade.NewUserAdminFacade(
		authHelperInst,
		userAdminGetUCInst,
		userAdminGetByNameUCInst,
		userAdminListUCInst,
		userAdminSaveUCInst,
		userAdminDeleteUCInst,
		confInst.App.MaxListLimit,
	), nil
}
