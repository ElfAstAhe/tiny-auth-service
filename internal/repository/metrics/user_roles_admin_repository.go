package metrics

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRolesAdminMetricsRepository implements libdomain.OwnedRepository and domain.UserRolesAdminRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component for administrative owned role entity storage operations.
// It intercepts data access vectors to observe and report execution performance states via the Prometheus telemetry engine.
type UserRolesAdminMetricsRepository struct {
	*metrics.BaseOwnedMetricsRepository[*domain.Role, string, string]                                 // Generic framework-level metrics base decorator structure for administrative owned aggregates
	repo                                                              domain.UserRolesAdminRepository // Downstream administrative data storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesAdminMetricsRepository)(nil)
var _ domain.UserRolesAdminRepository = (*UserRolesAdminMetricsRepository)(nil)

// NewUserRolesAdminMetricsRepository acts as a factory constructor mounting telemetry metrics collection proxies over administrative user-roles storage adapters.
func NewUserRolesAdminMetricsRepository(repo domain.UserRolesAdminRepository) *UserRolesAdminMetricsRepository {
	return &UserRolesAdminMetricsRepository{
		repo:                       repo,
		BaseOwnedMetricsRepository: metrics.NewBaseOwnedMetricsRepository[*domain.Role, string, string]("UserRolesAdminRepository", repo),
	}
}
