package repository

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserMetricsRepository struct {
	*repository.BaseCRUDMetricsRepository[*domain.User, string]

	repo domain.UserRepository
}

func NewUserMetricsRepository(repo domain.UserRepository) *UserMetricsRepository {
	return &UserMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: repository.NewBaseCRUDMetricsRepository[*domain.User, string]("UserRepository", repo),
	}
}

func (umr *UserMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.User, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(umr.BaseCRUDMetricsRepository.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return umr.repo.FindByName(ctx, name)
}
