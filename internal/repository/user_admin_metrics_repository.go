package repository

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserAdminMetricsRepository struct {
	*repository.BaseCRUDMetricsRepository[*domain.User, string]
	repo domain.UserAdminRepository
}

func NewUserAdminMetricsRepository(repo domain.UserAdminRepository) *UserAdminMetricsRepository {
	return &UserAdminMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: repository.NewBaseCRUDMetricsRepository[*domain.User, string]("UserAdminRepository", repo),
	}
}

func (uam *UserAdminMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(uam.BaseCRUDMetricsRepository.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return uam.repo.FindByName(ctx, name)
}
