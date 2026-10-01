package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceRoleRepo tracks the standard token key for primary public role storage database components.
	InstanceRoleRepo string = "roleRepository"
	// InstanceRoleMetricsRepo maps the telemetry metrics registration identifier targeting role persistence.
	InstanceRoleMetricsRepo string = "roleMetricsRepository"
	// InstanceRoleTraceRepo sets unique system keys identifying OpenTelemetry distributed tracing proxies over role repositories.
	InstanceRoleTraceRepo string = "roleTraceRepository"
	// InstanceRoleAuditRepo bridges structural identifiers executing database change audit logging over public roles.
	InstanceRoleAuditRepo string = "roleAuditRepository"

	// InstanceRoleAdminRepo tracks the standard token key for primary administrative role storage database components.
	InstanceRoleAdminRepo string = "roleAdminRepository"
	// InstanceRoleAdminMetricsRepo maps telemetry metrics identifiers targeting administrative role persistence proxies.
	InstanceRoleAdminMetricsRepo string = "roleAdminMetricsRepository"
	// InstanceRoleAdminTraceRepo sets unique system keys managing OpenTelemetry distributed tracing proxies over administrative role repositories.
	InstanceRoleAdminTraceRepo string = "roleAdminTraceRepository"
	// InstanceRoleAdminAuditRepo bridges structural identifiers executing database change audit logging over administrative roles.
	InstanceRoleAdminAuditRepo string = "roleAdminAuditRepository"

	// InstanceUserRepo tracks the standard token key for primary public user storage database components.
	InstanceUserRepo string = "userRepository"
	// InstanceUserMetricsRepo maps the telemetry metrics registration identifier targeting user persistence.
	InstanceUserMetricsRepo string = "userMetricsRepository"
	// InstanceUserTraceRepo sets unique system keys identifying OpenTelemetry distributed tracing proxies over user repositories.
	InstanceUserTraceRepo string = "userTraceRepository"
	// InstanceUserAuditRepo bridges structural identifiers executing database change audit logging over public users.
	InstanceUserAuditRepo string = "userAuditRepository"

	// InstanceUserAdminRepo tracks the standard token key for primary administrative user storage database components.
	InstanceUserAdminRepo string = "userAdminRepository"
	// InstanceUserAdminMetricsRepo maps telemetry metrics identifiers targeting administrative user persistence proxies.
	InstanceUserAdminMetricsRepo string = "userAdminMetricsRepository"
	// InstanceUserAdminTraceRepo sets unique system keys managing OpenTelemetry distributed tracing proxies over administrative user repositories.
	InstanceUserAdminTraceRepo string = "userAdminTraceRepository"
	// InstanceUserAdminAuditRepo bridges structural identifiers executing database change audit logging over administrative users.
	InstanceUserAdminAuditRepo string = "userAdminAuditRepository"

	// InstanceUserRolesRepo tracks lookup tokens mapping many-to-many bridge database components for public roles allocations.
	InstanceUserRolesRepo string = "userRolesRepository"
	// InstanceUserRolesMetricsRepo sets telemetry metric trackers over standard public many-to-many role link tables.
	InstanceUserRolesMetricsRepo string = "userRolesMetricsRepository"
	// InstanceUserRolesTraceRepo orchestrates OpenTelemetry spans across public relational bridge mapping operations.
	InstanceUserRolesTraceRepo string = "userRolesTraceRepository"

	// InstanceUserRolesAdminRepo tracks lookup tokens mapping many-to-many bridge database components for elevated administrative roles allocations.
	InstanceUserRolesAdminRepo string = "userRolesAdminRepository"
	// InstanceUserRolesAdminMetricsRepo sets telemetry metric trackers over elevated administrative many-to-many role link tables.
	InstanceUserRolesAdminMetricsRepo string = "userRolesAdminMetricsRepository"
	// InstanceUserRolesAdminTraceRepo orchestrates OpenTelemetry spans across administrative relational bridge mapping operations.
	InstanceUserRolesAdminTraceRepo string = "userRolesAdminTraceRepository"
)

// RepositoryContainer structures a lazy-loaded lifecycle dependency injection container managing storage layer data access objects and decorators chains.
type RepositoryContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*RepositoryContainer)(nil)
var _ container.LazyContainer = (*RepositoryContainer)(nil)

// NewRepositoryContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewRepositoryContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *RepositoryContainer {
	return &RepositoryContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(RepositoryContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking required lazy storage and decorator provider callbacks inside the dependency graph.
func (rc *RepositoryContainer) Init(ctx context.Context) error {
	err := errors.Join(
		rc.RegisterProvider(InstanceRoleRepo, rc.providerRoleRepo),
		rc.RegisterProvider(InstanceRoleMetricsRepo, rc.providerRoleMetricsRepo),
		rc.RegisterProvider(InstanceRoleTraceRepo, rc.providerRoleTraceRepo),
		rc.RegisterProvider(InstanceRoleAuditRepo, rc.providerRoleAuditRepo),

		rc.RegisterProvider(InstanceRoleAdminRepo, rc.providerRoleAdminRepo),
		rc.RegisterProvider(InstanceRoleAdminMetricsRepo, rc.providerRoleAdminMetricsRepo),
		rc.RegisterProvider(InstanceRoleAdminTraceRepo, rc.providerRoleAdminTraceRepo),
		rc.RegisterProvider(InstanceRoleAdminAuditRepo, rc.providerRoleAdminAuditRepo),

		rc.RegisterProvider(InstanceUserRepo, rc.providerUserRepo),
		rc.RegisterProvider(InstanceUserMetricsRepo, rc.providerUserMetricsRepo),
		rc.RegisterProvider(InstanceUserTraceRepo, rc.providerUserTraceRepo),
		rc.RegisterProvider(InstanceUserAuditRepo, rc.providerUserAuditRepo),

		rc.RegisterProvider(InstanceUserAdminRepo, rc.providerUserAdminRepo),
		rc.RegisterProvider(InstanceUserAdminMetricsRepo, rc.providerUserAdminMetricsRepo),
		rc.RegisterProvider(InstanceUserAdminTraceRepo, rc.providerUserAdminTraceRepo),
		rc.RegisterProvider(InstanceUserAdminAuditRepo, rc.providerUserAdminAuditRepo),

		rc.RegisterProvider(InstanceUserRolesRepo, rc.providerUserRolesRepo),
		rc.RegisterProvider(InstanceUserRolesMetricsRepo, rc.providerUserRolesMetricsRepo),
		rc.RegisterProvider(InstanceUserRolesTraceRepo, rc.providerUserRolesTraceRepo),

		rc.RegisterProvider(InstanceUserRolesAdminRepo, rc.providerUserRolesAdminRepo),
		rc.RegisterProvider(InstanceUserRolesAdminMetricsRepo, rc.providerUserRolesAdminMetricsRepo),
		rc.RegisterProvider(InstanceUserRolesAdminTraceRepo, rc.providerUserRolesAdminTraceRepo),
	)
	if err != nil {
		return errs.NewContainerError(rc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
