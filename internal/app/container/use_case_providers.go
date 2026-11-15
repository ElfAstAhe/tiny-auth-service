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
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase/trace"
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

	return trace.NewChangeKeysUseCase(
		"ChangeKeysUseCase",
		usecase.NewChangeKeys(
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

	return trace.NewChangePasswordUseCase(
		"ChangePasswordUseCase",
		usecase.NewChangePassword(
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

	return trace.NewLoginTraceUseCase(
		"LoginUseCase",
		usecase.NewLoginCiphered(
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

	return trace.NewLoginBasicUseCase(
		"LoginSimpleUseCase",
		usecase.NewLoginBasic(
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

	return trace.NewProfileTraceUseCase("ProfileUseCase", usecase.NewProfile(userRepoInst)), nil
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

	return trace.NewRegisterUseCase(
		"RegisterUseCase",
		usecase.NewRegister(
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

	return trace.NewRoleAdminDeleteUseCase(
		"RoleAdminDeleteUseCase",
		usecase.NewRoleAdminDelete(uwInst, roleAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminGetUC() (any, error) {
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return trace.NewRoleAdminGetUseCase("RoleAdminGetUseCase", usecase.NewRoleAdminGet(roleAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerRoleAdminGetByNameUC() (any, error) {
	roleAdminRepoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return trace.NewRoleAdminGetByNameUseCase("RoleAdminGetByNameUseCase", usecase.NewRoleAdminGetByName(roleAdminRepoInst)), nil
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

	return trace.NewRoleAdminListUseCase(
		"RoleAdminListUseCase",
		usecase.NewRoleAdminList(roleAdminRepoInst, confInst.App.MaxListLimit)), nil
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

	return trace.NewRoleAdminSaveUseCase(
		"RoleAdminSaveUseCase",
		usecase.NewRoleAdminSave(uwInst, roleAdminRepoInst)), nil
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

	return trace.NewUserAdminDeleteUseCase(
		"UserAdminDeleteUseCase",
		usecase.NewUserAdminDelete(uwInst, userAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminGetUC() (any, error) {
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return trace.NewUserAdminGetUseCase(
		"UserAdminGetUseCase",
		usecase.NewUserAdminGet(userAdminRepoInst)), nil
}

//goland:noinspection DuplicatedCode
func (ucc *UseCaseContainer) providerUserAdminGetByNameUC() (any, error) {
	userAdminRepoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminAuditRepo)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return trace.NewUserAdminGetByNameUseCase(
		"UserAdminGetByNameUseCase",
		usecase.NewUserAdminGetByName(userAdminRepoInst)), nil
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

	return trace.NewUserAdminListUseCase(
		"UserAdminListUseCase",
		usecase.NewUserAdminList(userAdminRepoInst, confInst.App.MaxListLimit)), nil
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

	return trace.NewUserAdminSaveUseCase(
		"UserAdminSaveUseCase",
		usecase.NewUserAdminSave(
			uwInst,
			hashCipherInst,
			keysHelperInst,
			userAdminRepoInst,
		)), nil
}
