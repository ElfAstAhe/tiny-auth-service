package repository

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type RoleAdminMetricsRepository struct {
	*repository.BaseMetricsRepository[*domain.Role, string]

	repo domain.RoleRepository
}

func NewRoleAdminMetricsRepository(repo domain.RoleAdminRepository) *RoleAdminMetricsRepository {
	return &RoleAdminMetricsRepository{
		repo:                  repo,
		BaseMetricsRepository: repository.NewBaseMetricsRepository[*domain.Role, string](repo),
	}
}

func (ram *RoleAdminMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.Role, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(ram.BaseMetricsRepository.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return ram.repo.FindByName(ctx, name)
}
