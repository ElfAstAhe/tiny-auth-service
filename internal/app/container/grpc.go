package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceAuthGRPCService tracks the registry descriptor key for the baseline authentication gRPC transport handler.
	InstanceAuthGRPCService string = "auth-gRPC-service"
	// InstanceUserGRPCService maps the unique lookup token for the core user profiles query gRPC endpoint.
	InstanceUserGRPCService string = "user-gRPC-service"
	// InstanceUserAdminGRPCService pins the administrative identity moderation gRPC delivery interactor wrapper.
	InstanceUserAdminGRPCService string = "user-admin-gRPC-service"
	// InstanceRoleAdminGRPCService setups lookup tags routing role access matrices modifiers commands to gRPC listeners.
	InstanceRoleAdminGRPCService string = "role-admin-gRPC-service"
	// InstanceGRPCRunner structures the core library server pipeline manager running networking listening threads loops.
	InstanceGRPCRunner string = "grpc-runner"
)

// GRPCContainer structures a lazy-loaded lifecycle dependency injection container managing high-performance binary Protobuf transport components.
type GRPCContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*GRPCContainer)(nil)
var _ container.LazyContainer = (*GRPCContainer)(nil)

// NewGRPCContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewGRPCContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *GRPCContainer {
	return &GRPCContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(GRPCContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking binary protobuf services and server runners callbacks inside the dependency graph.
func (gc *GRPCContainer) Init(ctx context.Context) error {
	err := errors.Join(
		gc.RegisterProvider(InstanceAuthGRPCService, gc.providerAuthService),
		gc.RegisterProvider(InstanceUserGRPCService, gc.providerUserService),
		gc.RegisterProvider(InstanceUserAdminGRPCService, gc.providerUserAdminService),
		gc.RegisterProvider(InstanceRoleAdminGRPCService, gc.providerRoleAdminService),
		gc.RegisterProvider(InstanceGRPCRunner, gc.providerGRPCRunner),
	)
	if err != nil {
		return errs.NewContainerError(gc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
