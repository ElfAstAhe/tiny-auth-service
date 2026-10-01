package container

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	libuc "github.com/ElfAstAhe/go-service-template/pkg/usecase"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase/telemetry"
)

func (ucc *UseCaseContainer) providerUnitOfWork() (any, error) {
	tmInst, err := container.GetInstance[db.TransactionManager](InstanceTM)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return libuc.NewUnitOfWork(tmInst, nil), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerChangeKeysUC() (any, error) {
	keysHelperInst, err := container.GetInstance[helper.RSAKeys](InstanceKeysHelper)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userRepoInst, err := container.GetInstance[domain.UserRepository](InstanceUserAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewChangeKeysTraceUseCase(
		"ChangeKeysUseCase",
		usecase.NewChangeKeysUseCase(
			keysHelperInst,
			uwInst,
			userRepoInst,
		)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerChangePasswordUC() (any, error) {
	hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userRepoInst, err := container.GetInstance[domain.UserRepository](InstanceUserAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewChangePasswordTraceUseCase(
		"ChangePasswordUseCase",
		usecase.NewChangePasswordUseCase(
			hashCipherInst,
			uwInst,
			userRepoInst,
		)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerLoginUC() (any, error) {
	hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	keysHelperInst, err := container.GetInstance[helper.RSAKeys](InstanceKeysHelper)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userRepoInst, err := container.GetInstance[domain.UserRepository](InstanceUserAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewLoginTraceUseCase(
		"LoginUseCase",
		usecase.NewLoginUseCase(
			hashCipherInst,
			keysHelperInst,
			authHelperInst,
			userRepoInst,
		)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerLoginSimpleUC() (any, error) {
	hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userRepoInst, err := container.GetInstance[domain.UserRepository](InstanceUserAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewLoginSimpleTraceUseCase(
		"LoginSimpleUseCase",
		usecase.NewLoginSimpleUseCase(
			hashCipherInst,
			authHelperInst,
			userRepoInst,
		)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerProfileUC() (any, error) {
	userRepoInst, err := container.GetInstance[domain.UserRepository](InstanceUserAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewProfileTraceUseCase("ProfileUseCase", usecase.NewProfileUseCase(userRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRegisterUC() (any, error) {
	hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	keysHelperInst, err := container.GetInstance[helper.RSAKeys](InstanceKeysHelper)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userRepoInst, err := container.GetInstance[domain.UserRepository](InstanceUserAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewRegisterTraceUseCase(
		"RegisterUseCase",
		usecase.NewRegisterUseCase(
			uwInst,
			hashCipherInst,
			keysHelperInst,
			userRepoInst,
		)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminDeleteUC() (any, error) {
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewRoleAdminDeleteTraceUseCase(
		"RoleAdminDeleteUseCase",
		usecase.NewRoleAdminDeleteUseCase(uwInst, roleAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminGetUC() (any, error) {
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewRoleAdminGetTraceUseCase("RoleAdminGetUseCase", usecase.NewRoleAdminGetUseCase(roleAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminGetByNameUC() (any, error) {
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewRoleAdminGetNameTraceUseCase("RoleAdminGetByNameUseCase", usecase.NewRoleAdminGetNameUseCase(roleAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminListUC() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewRoleAdminListTraceUseCase(
		"RoleAdminListUseCase",
		usecase.NewRoleAdminListUseCase(roleAdminRepoInst, confInst.App.MaxListLimit)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminSaveUC() (any, error) {
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewRoleAdminSaveTraceUseCase(
		"RoleAdminSaveUseCase",
		usecase.NewRoleAdminSaveUseCase(uwInst, roleAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminDeleteUC() (any, error) {
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewUserAdminDeleteTraceUseCase(
		"UserAdminDeleteUseCase",
		usecase.NewUserAdminDeleteUseCase(uwInst, userAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminGetUC() (any, error) {
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewUserAdminGetTraceUseCase(
		"UserAdminGetUseCase",
		usecase.NewUserAdminGetUseCase(userAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminGetByNameUC() (any, error) {
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewUserAdminGetNameTraceUseCase(
		"UserAdminGetByNameUseCase",
		usecase.NewUserAdminGetNameUseCase(userAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminListUC() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewUserAdminListTraceUseCase(
		"UserAdminListUseCase",
		usecase.NewUserAdminListUseCase(userAdminRepoInst, confInst.App.MaxListLimit)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminSaveUC() (any, error) {
	hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	keysHelperInst, err := container.GetInstance[helper.RSAKeys](InstanceKeysHelper)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	uwInst, err := container.GetInstance[libdom.UnitOfWork](InstanceUnitOfWork)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return telemetry.NewUserAdminSaveTraceUseCase(
		"UserAdminSaveUseCase",
		usecase.NewUserAdminSaveUseCase(
			uwInst,
			hashCipherInst,
			keysHelperInst,
			userAdminRepoInst,
		)), nil
}
