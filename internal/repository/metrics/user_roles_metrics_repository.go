package metrics

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesMetricsRepository struct {
	*repository.BaseOwnedMetricsRepository[*domain.Role, string, string]
	repo domain.UserRolesRepository
}

var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesMetricsRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesMetricsRepository)(nil)

func NewUserRolesMetricsRepository(repo domain.UserRolesRepository) *UserRolesMetricsRepository {
	return &UserRolesMetricsRepository{
		repo:                       repo,
		BaseOwnedMetricsRepository: repository.NewBaseOwnedMetricsRepository[*domain.Role, string, string]("UserRolesRepository", repo),
	}
}
