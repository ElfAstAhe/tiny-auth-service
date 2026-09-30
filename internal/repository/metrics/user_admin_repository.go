package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserAdminMetricsRepository struct {
	*metrics.BaseCRUDMetricsRepository[*domain.User, string]
	repo domain.UserAdminRepository
}

var _ libdomain.CRUDRepository[*domain.User, string] = (*UserAdminMetricsRepository)(nil)
var _ domain.UserAdminRepository = (*UserAdminMetricsRepository)(nil)

func NewUserAdminMetricsRepository(repo domain.UserAdminRepository) *UserAdminMetricsRepository {
	return &UserAdminMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: metrics.NewBaseCRUDMetricsRepository[*domain.User, string]("UserAdminRepository", repo),
	}
}

func (uam *UserAdminMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(uam.BaseCRUDMetricsRepository.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return uam.repo.FindByName(ctx, name)
}
