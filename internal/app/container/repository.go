package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

const (
	InstanceRoleRepo        string = "roleRepository"
	InstanceRoleMetricsRepo string = "roleMetricsRepository"
	InstanceRoleTraceRepo   string = "roleTraceRepository"
	InstanceRoleAuditRepo   string = "roleAuditRepository"

	InstanceRoleAdminRepo        string = "roleAdminRepository"
	InstanceRoleAdminMetricsRepo string = "roleAdminMetricsRepository"
	InstanceRoleAdminTraceRepo   string = "roleAdminTraceRepository"
	InstanceRoleAdminAuditRepo   string = "roleAdminAuditRepository"

	InstanceUserRepo        string = "userRepository"
	InstanceUserMetricsRepo string = "userMetricsRepository"
	InstanceUserTraceRepo   string = "userTraceRepository"
	InstanceUserAuditRepo   string = "userAuditRepository"

	InstanceUserAdminRepo        string = "userAdminRepository"
	InstanceUserAdminMetricsRepo string = "userAdminMetricsRepository"
	InstanceUserAdminTraceRepo   string = "userAdminTraceRepository"
	InstanceUserAdminAuditRepo   string = "userAdminAuditRepository"

	InstanceUserRolesRepo        string = "userRolesRepository"
	InstanceUserRolesMetricsRepo string = "userRolesMetricsRepository"
	InstanceUserRolesTraceRepo   string = "userRolesTraceRepository"

	InstanceUserRolesAdminRepo        string = "userRolesAdminRepository"
	InstanceUserRolesAdminMetricsRepo string = "userRolesAdminMetricsRepository"
	InstanceUserRolesAdminTraceRepo   string = "userRolesAdminTraceRepository"
)

type RepositoryContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*RepositoryContainer)(nil)
var _ container.LazyContainer = (*RepositoryContainer)(nil)

func NewRepositoryContainer(orchestrator container.Orchestrator) *RepositoryContainer {
	return &RepositoryContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(RepositoryContainerName, orchestrator),
	}
}

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
