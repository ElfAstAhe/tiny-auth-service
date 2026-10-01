package metrics

import (
	"context"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserMetricsRepository implements libdom.CRUDRepository and domain.UserRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component.
// It wraps data access methods to record performance indicators and latency histograms using the Prometheus engine.
type UserMetricsRepository struct {
	*metrics.BaseCRUDMetricsRepository[*domain.User, string] // Generic framework-level metrics base decorator structure

	repo domain.UserRepository // Downstream data storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdom.CRUDRepository[*domain.User, string] = (*UserMetricsRepository)(nil)
var _ domain.UserRepository = (*UserMetricsRepository)(nil)

// NewUserMetricsRepository acts as a factory constructor mounting telemetry metric proxies over user storage adapters.
func NewUserMetricsRepository(repo domain.UserRepository) *UserMetricsRepository {
	return &UserMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: metrics.NewBaseCRUDMetricsRepository[*domain.User, string]("UserRepository", repo),
	}
}

// FindByName proxies user lookup sequences, automatically capturing operational durations and reporting execution states to collectors.
func (umr *UserMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(umr.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return umr.repo.FindByName(ctx, name)
}
