package metrics

import (
	"context"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRepository implements libdom.CRUDRepository and domain.UserRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component.
// It wraps data access methods to record performance indicators and latency histograms using the Prometheus engine.
type UserRepository struct {
	*metrics.CRUD[*domain.User, string] // Generic framework-level metrics base decorator structure

	repo domain.UserRepository // Downstream data storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdom.CRUDRepository[*domain.User, string] = (*UserRepository)(nil)
var _ domain.UserRepository = (*UserRepository)(nil)

// NewUserRepository acts as a factory constructor mounting telemetry metric proxies over user storage adapters.
func NewUserRepository(repo domain.UserRepository) *UserRepository {
	return &UserRepository{
		repo: repo,
		CRUD: metrics.NewCRUD[*domain.User, string]("UserRepository", repo),
	}
}

// FindByName proxies user lookup sequences, automatically capturing operational durations and reporting execution states to collectors.
func (umr *UserRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(umr.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return umr.repo.FindByName(ctx, name)
}
