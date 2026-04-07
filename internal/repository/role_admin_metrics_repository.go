package repository

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type RoleAdminMetricsRepository struct {
	*repository.BaseCRUDMetricsRepository[*domain.Role, string]

	repo domain.RoleRepository
}

var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleAdminMetricsRepository)(nil)
var _ domain.RoleAdminRepository = (*RoleAdminMetricsRepository)(nil)

func NewRoleAdminMetricsRepository(repo domain.RoleAdminRepository) *RoleAdminMetricsRepository {
	return &RoleAdminMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: repository.NewBaseCRUDMetricsRepository[*domain.Role, string]("RoleAdminRepository", repo),
	}
}

func (ram *RoleAdminMetricsRepository) FindByName(ctx context.Context, name string) (res *domain.Role, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(ram.BaseCRUDMetricsRepository.GetRepositoryName(), "FindByName", err, start)
	}(time.Now())

	return ram.repo.FindByName(ctx, name)
}
