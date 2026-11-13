package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminRepository implements libdomain.CRUDRepository and domain.UserAdminRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component for administrative user storage operations.
// It wraps data access methods to record performance indicators and latency histograms using the Prometheus engine.
type UserAdminRepository struct {
	*metrics.CRUD[*domain.User, string]                            // Generic framework-level metrics base decorator structure
	repo                                domain.UserAdminRepository // Downstream core administrative database persistence adapter layer
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.User, string] = (*UserAdminRepository)(nil)
var _ domain.UserAdminRepository = (*UserAdminRepository)(nil)

// NewUserAdminRepository acts as a factory constructor mounting telemetry metric proxies over administrative user storage adapters.
func NewUserAdminRepository(repo domain.UserAdminRepository) *UserAdminRepository {
	return &UserAdminRepository{
		repo: repo,
		CRUD: metrics.NewCRUD[*domain.User, string]("UserAdminRepository", repo),
	}
}

// FindByName proxies user lookup sequences, automatically capturing operational durations and reporting execution states to collectors.
func (uam *UserAdminRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(uam.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return uam.repo.FindByName(ctx, name)
}
