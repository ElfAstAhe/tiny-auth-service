package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceUnitOfWork defines the lookup token key for the global database atomic transaction manager boundary boundary.
	InstanceUnitOfWork string = "unit-of-work"

	// InstanceChangeKeysUC tracks the registration identifier for asymmetric key pair rotation business interactor.
	InstanceChangeKeysUC string = "change-keys-uc"
	// InstanceChangePasswordUC maps the configuration identifier for secure profile password modification workflow interactor.
	InstanceChangePasswordUC string = "change-password-uc"
	// InstanceLoginUC sets the lookup token mapping full cryptographic security validation login usecases boundaries.
	InstanceLoginUC string = "login-uc"
	// InstanceLoginSimpleUC tracks identifiers servicing lightweight machine credentials verification procedures workflows.
	InstanceLoginSimpleUC string = "login-simple-uc"
	// InstanceProfileUC links targets resolving structural non-administrative profile properties queries scripts.
	InstanceProfileUC string = "profile-uc"
	// InstanceRegisterUC sets unique system keys identifying registration and asset provisioning workflows.
	InstanceRegisterUC string = "register-uc"

	// InstanceRoleAdminDeleteUC maps identifiers running administrative role eviction operations commands down.
	InstanceRoleAdminDeleteUC string = "role-admin-delete-uc"
	// InstanceRoleAdminGetUC defines unique lookup tags for specific identifier roles retrieval requests execution.
	InstanceRoleAdminGetUC string = "role-admin-get-uc"
	// InstanceRoleAdminGetByNameUC maps token handles matching specific target role name text lookup parameters queries.
	InstanceRoleAdminGetByNameUC string = "role-admin-get-by-name-uc"
	// InstanceRoleAdminListUC structures identifier boundaries allocated to paginated role collection aggregate arrays extractions.
	InstanceRoleAdminListUC string = "role-admin-list-uc"
	// InstanceRoleAdminSaveUC links targets running administrative role modification and creation persistence workflows.
	InstanceRoleAdminSaveUC string = "role-admin-save-uc"

	// InstanceUserAdminDeleteUC structures specific keys mapping destructive user moderator eviction parameters boundaries.
	InstanceUserAdminDeleteUC string = "user-admin-delete-uc"
	// InstanceUserAdminGetUC pins unique registry passport descriptors dedicated to single user metadata lookups.
	InstanceUserAdminGetUC string = "user-admin-get-uc"
	// InstanceUserAdminGetByNameUC sets string keys managing administrative user profile name exact phrase queries search options.
	InstanceUserAdminGetByNameUC string = "user-admin-get-by-name-uc"
	// InstanceUserAdminListUC tracks boundaries allocated for paginated collection arrays extraction targeting registered user profiles.
	InstanceUserAdminListUC string = "user-admin-list-uc"
	// InstanceUserAdminSaveUC orchestrates required target identifiers execution routes balancing user profiles updates and insertion commands.
	InstanceUserAdminSaveUC string = "user-admin-save-uc"
)

// UseCaseContainer structures a lazy-loaded lifecycle dependency injection container managing application business logic interactors.
type UseCaseContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*UseCaseContainer)(nil)
var _ container.LazyContainer = (*UseCaseContainer)(nil)

// NewUseCaseContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
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

// Init triggers synchronous registration sequences linking required lazy interactor provider callbacks inside the dependency map graph.
func (ucc *UseCaseContainer) Init(ctx context.Context) error {
	err := errors.Join(
		ucc.RegisterProvider(InstanceUnitOfWork, ucc.providerUnitOfWork),

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
