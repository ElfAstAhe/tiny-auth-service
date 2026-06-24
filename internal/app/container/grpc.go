package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	InstanceAuthGRPCService      string = "auth-gRPC-service"
	InstanceUserGRPCService      string = "user-gRPC-service"
	InstanceUserAdminGRPCService string = "user-admin-gRPC-service"
	InstanceRoleAdminGRPCService string = "role-admin-gRPC-service"
	InstanceGRPCRunner           string = "grpc-runner"
)

type GRPCContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*GRPCContainer)(nil)
var _ container.LazyContainer = (*GRPCContainer)(nil)

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
