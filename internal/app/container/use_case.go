package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	InstanceTM string = "TransactionManager"

	InstanceChangeKeysUC     string = "change-keys-uc"
	InstanceChangePasswordUC string = "change-password-uc"
	InstanceLoginUC          string = "login-uc"
	InstanceLoginSimpleUC    string = "login-simple-uc"
	InstanceProfileUC        string = "profile-uc"
	InstanceRegisterUC       string = "register-uc"

	InstanceRoleAdminDeleteUC    string = "role-admin-delete-uc"
	InstanceRoleAdminGetUC       string = "role-admin-get-uc"
	InstanceRoleAdminGetByNameUC string = "role-admin-get-by-name-uc"
	InstanceRoleAdminListUC      string = "role-admin-list-uc"
	InstanceRoleAdminSaveUC      string = "role-admin-save-uc"

	InstanceUserAdminDeleteUC    string = "user-admin-delete-uc"
	InstanceUserAdminGetUC       string = "user-admin-get-uc"
	InstanceUserAdminGetByNameUC string = "user-admin-get-by-name-uc"
	InstanceUserAdminListUC      string = "user-admin-list-uc"
	InstanceUserAdminSaveUC      string = "user-admin-save-uc"
)

type UseCaseContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*UseCaseContainer)(nil)
var _ container.LazyContainer = (*UseCaseContainer)(nil)

func NewUseCaseContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *UseCaseContainer {
	return &UseCaseContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(UseCaseContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

func (ucc *UseCaseContainer) Init(ctx context.Context) error {
	err := errors.Join(
		ucc.RegisterProvider(InstanceTM, ucc.providerTM),

		ucc.RegisterProvider(InstanceChangeKeysUC, ucc.providerChangeKeysUC),
		ucc.RegisterProvider(InstanceChangePasswordUC, ucc.providerChangePasswordUC),
		ucc.RegisterProvider(InstanceLoginUC, ucc.providerLoginUC),
		ucc.RegisterProvider(InstanceLoginSimpleUC, ucc.providerLoginSimpleUC),
		ucc.RegisterProvider(InstanceProfileUC, ucc.providerProfileUC),
		ucc.RegisterProvider(InstanceRegisterUC, ucc.providerRegisterUC),

		ucc.RegisterProvider(InstanceRoleAdminDeleteUC, ucc.providerRoleAdminDeleteUC),
		ucc.RegisterProvider(InstanceRoleAdminGetUC, ucc.providerRoleAdminGetUC),
		ucc.RegisterProvider(InstanceRoleAdminGetByNameUC, ucc.providerRoleAdminGetByNameUC),
		ucc.RegisterProvider(InstanceRoleAdminListUC, ucc.providerRoleAdminListUC),
		ucc.RegisterProvider(InstanceRoleAdminSaveUC, ucc.providerRoleAdminSaveUC),

		ucc.RegisterProvider(InstanceUserAdminDeleteUC, ucc.providerUserAdminDeleteUC),
		ucc.RegisterProvider(InstanceUserAdminGetUC, ucc.providerUserAdminGetUC),
		ucc.RegisterProvider(InstanceUserAdminGetByNameUC, ucc.providerUserAdminGetByNameUC),
		ucc.RegisterProvider(InstanceUserAdminListUC, ucc.providerUserAdminListUC),
		ucc.RegisterProvider(InstanceUserAdminSaveUC, ucc.providerUserAdminSaveUC),
	)
	if err != nil {
		return errs.NewContainerError(ucc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
