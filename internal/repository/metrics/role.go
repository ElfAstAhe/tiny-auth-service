package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleRepository implements libdomain.CRUDRepository and domain.RoleRepository interfaces,
// acting as a non-invasive telemetry metrics decorator component for role storage operations.
// It wraps data access methods to record performance indicators and latency histograms using the Prometheus engine.
type RoleRepository struct {
	*metrics.CRUD[*domain.Role, string]                       // Generic framework-level metrics base decorator structure
	repo                                domain.RoleRepository // Downstream data storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleRepository)(nil)
var _ domain.RoleRepository = (*RoleRepository)(nil)

// NewRoleRepository acts as a factory constructor mounting telemetry metric proxies over role storage adapters.
func NewRoleRepository(repo domain.RoleRepository) *RoleRepository {
	return &RoleRepository{
		repo: repo,
		CRUD: metrics.NewCRUD[*domain.Role, string]("RoleRepository", repo),
	}
}

// FindByName proxies role lookup sequences, automatically capturing operational durations and reporting execution states to collectors.
func (rmr *RoleRepository) FindByName(ctx context.Context, name string) (res *domain.Role, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(rmr.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return rmr.repo.FindByName(ctx, name)
}
