package container

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceApplication tracks the registration lookup key for the root system lifecycle engine.
	InstanceApplication string = "application"
	// InstanceApplicationReady defines the key descriptor mapping readiness synchronization handles.
	InstanceApplicationReady string = "application-ready"
	// InstanceApplicationHealth bridges core infrastructure health monitoring nodes.
	InstanceApplicationHealth string = "application-health"
	// InstanceConfig sets unique system keys identifying global configuration blueprints.
	InstanceConfig string = "config"
	// InstanceLogger targets the shared platform diagnostics logging component.
	InstanceLogger string = "logger"
)

// AppContainer structures the non-lazy baseline application container orchestrating foundational runtime configurations.
type AppContainer struct {
	*container.BaseContainer // Generic framework-level root baseline container handle
}

// Compile-time interface compliance verification
var _ container.Container = (*AppContainer)(nil)

// NewAppContainer acts as a factory constructor deploying core application container names and logging boundaries.
func NewAppContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *AppContainer {
	return &AppContainer{
		BaseContainer: container.NewBaseContainer(
			container.WithName(AppContainerName),
			container.WithOrchestrator(orchestrator),
			container.WithLogger(log),
		),
	}
}

// Init handles early lifecycle initialization steps, acting as a placeholder loop within the static graph.
func (ac *AppContainer) Init(ctx context.Context) error {
	return nil
}

// Close ensures proper resource evictions upon active server termination procedures.
func (ac *AppContainer) Close(ctx context.Context) error {
	return nil
}
