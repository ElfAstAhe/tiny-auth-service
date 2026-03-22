package repository

import (
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesMetricsRepository struct {
	*repository.BaseOwnedMetricsRepository[*domain.Role, string, string]
	repo domain.UserRolesRepository
}

func NewUserRolesMetricsRepository(repo domain.UserRolesRepository) *UserRolesMetricsRepository {
	return &UserRolesMetricsRepository{
		repo:                       repo,
		BaseOwnedMetricsRepository: repository.NewBaseOwnedMetricsRepository[*domain.Role, string, string]("UserRolesRepository", repo),
	}
}
