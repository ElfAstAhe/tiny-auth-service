package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type RoleMetricsRepository struct {
	*repository.BaseCRUDMetricsRepository[*domain.Role, string]
	repo domain.RoleRepository
}

var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleMetricsRepository)(nil)
var _ domain.RoleRepository = (*RoleMetricsRepository)(nil)

func NewRoleMetricsRepository(repo domain.RoleRepository) *RoleMetricsRepository {
	return &RoleMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: repository.NewBaseCRUDMetricsRepository[*domain.Role, string]("RoleRepository", repo),
	}
}

func (rmr *RoleMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.Role, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(rmr.BaseCRUDMetricsRepository.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return rmr.repo.FindByName(ctx, name)
}
