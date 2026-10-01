package metrics

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRolesMetricsRepository implements libdomain.OwnedRepository and domain.UserRolesRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component for owned role entity storage operations.
// It intercepts data access vectors to observe and report execution performance states via the Prometheus telemetry engine.
type UserRolesMetricsRepository struct {
	*metrics.BaseOwnedMetricsRepository[*domain.Role, string, string]                            // Generic framework-level metrics base decorator structure for owned aggregates
	repo                                                              domain.UserRolesRepository // Downstream data storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesMetricsRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesMetricsRepository)(nil)

// NewUserRolesMetricsRepository acts as a factory constructor mounting telemetry metrics collection proxies over user-roles storage adapters.
func NewUserRolesMetricsRepository(repo domain.UserRolesRepository) *UserRolesMetricsRepository {
	return &UserRolesMetricsRepository{
		repo:                       repo,
		BaseOwnedMetricsRepository: metrics.NewBaseOwnedMetricsRepository[*domain.Role, string, string]("UserRolesRepository", repo),
	}
}
