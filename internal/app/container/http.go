package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceHTTPRouter defines the lookup token key targeted for registering the HTTP multiplexer router.
	InstanceHTTPRouter string = "HTTPRouter"
	// InstanceHTTPRunner structures the registration identifier for the network listener runner engine.
	InstanceHTTPRunner string = "HTTPRunner"
)

// HTTPContainer structures a lazy-loaded lifecycle dependency injection container managing HTTP transport and routing components.
type HTTPContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*HTTPContainer)(nil)
var _ container.LazyContainer = (*HTTPContainer)(nil)

// NewHTTPContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewHTTPContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *HTTPContainer {
	return &HTTPContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(HTTPContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking required lazy multiplexer router and server runner callbacks inside the dependency map graph.
func (hc *HTTPContainer) Init(initCtx context.Context) error {
	err := errors.Join(
		hc.RegisterProvider(InstanceHTTPRouter, hc.providerChiRouter),
		hc.RegisterProvider(InstanceHTTPRunner, hc.providerHTTPRunner),
	)
	if err != nil {
		return errs.NewContainerError(hc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
