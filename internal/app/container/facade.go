package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceAuthFacade tracks the global framework registration lookup token key targeted for session validation facade routers.
	InstanceAuthFacade string = "AuthFacade"
	// InstanceRoleAdminFacade maps the unique lookup token for administrative role management facade components.
	InstanceRoleAdminFacade string = "RoleAdminFacade"
	// InstanceUserFacade pins the public profile query and identity mutation facade boundary.
	InstanceUserFacade string = "UserFacade"
	// InstanceUserAdminFacade structures unique registry descriptors dedicated to administrative user moderation facades.
	InstanceUserAdminFacade string = "UserAdminFacade"
)

// FacadeContainer structures a lazy-loaded lifecycle dependency injection container managing application boundary facades routers.
type FacadeContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*FacadeContainer)(nil)
var _ container.LazyContainer = (*FacadeContainer)(nil)

// NewFacadeContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewFacadeContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *FacadeContainer {
	return &FacadeContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(FacadeContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking required lazy application boundary facade provider callbacks inside the graph.
func (fc *FacadeContainer) Init(ctx context.Context) error {
	err := errors.Join(
		fc.RegisterProvider(InstanceAuthFacade, fc.providerAuthFacade),
		fc.RegisterProvider(InstanceRoleAdminFacade, fc.providerRoleAdminFacade),
		fc.RegisterProvider(InstanceUserFacade, fc.providerUserFacade),
		fc.RegisterProvider(InstanceUserAdminFacade, fc.providerUserAdminFacade),
	)
	if err != nil {
		return errs.NewContainerError(fc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
