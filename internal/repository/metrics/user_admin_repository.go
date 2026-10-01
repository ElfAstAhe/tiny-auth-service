package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminMetricsRepository implements libdomain.CRUDRepository and domain.UserAdminRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component for administrative user storage operations.
// It wraps data access methods to record performance indicators and latency histograms using the Prometheus engine.
type UserAdminMetricsRepository struct {
	*metrics.BaseCRUDMetricsRepository[*domain.User, string]                            // Generic framework-level metrics base decorator structure
	repo                                                     domain.UserAdminRepository // Downstream core administrative database persistence adapter layer
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.User, string] = (*UserAdminMetricsRepository)(nil)
var _ domain.UserAdminRepository = (*UserAdminMetricsRepository)(nil)

// NewUserAdminMetricsRepository acts as a factory constructor mounting telemetry metric proxies over administrative user storage adapters.
func NewUserAdminMetricsRepository(repo domain.UserAdminRepository) *UserAdminMetricsRepository {
	return &UserAdminMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: metrics.NewBaseCRUDMetricsRepository[*domain.User, string]("UserAdminRepository", repo),
	}
}

// FindByName proxies user lookup sequences, automatically capturing operational durations and reporting execution states to collectors.
func (uam *UserAdminMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(uam.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return uam.repo.FindByName(ctx, name)
}
