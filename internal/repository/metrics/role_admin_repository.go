package metrics

import (
	"context"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminMetricsRepository implements libdomain.CRUDRepository and domain.RoleAdminRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component for administrative role storage operations.
// It wraps data access methods to record performance indicators and latency histograms using the Prometheus engine.
type RoleAdminMetricsRepository struct {
	*metrics.BaseCRUDMetricsRepository[*domain.Role, string] // Generic framework-level metrics base decorator structure

	repo domain.RoleAdminRepository // Downstream core administrative database persistence adapter layer
}

// Compile-time interface compliance verifications
var _ libdom.CRUDRepository[*domain.Role, string] = (*RoleAdminMetricsRepository)(nil)
var _ domain.RoleAdminRepository = (*RoleAdminMetricsRepository)(nil)

// NewRoleAdminMetricsRepository acts as a factory constructor mounting telemetry metric proxies over administrative role storage adapters.
func NewRoleAdminMetricsRepository(repo domain.RoleAdminRepository) *RoleAdminMetricsRepository {
	return &RoleAdminMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: metrics.NewBaseCRUDMetricsRepository[*domain.Role, string]("RoleAdminRepository", repo),
	}
}

// FindByName proxies administrative role lookup sequences, automatically capturing operational durations and reporting execution states to collectors.
func (ram *RoleAdminMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.Role, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(ram.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return ram.repo.FindByName(ctx, name)
}
