package metrics

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesAdminMetricsRepository struct {
	*metrics.BaseOwnedMetricsRepository[*domain.Role, string, string]
	repo domain.UserRolesAdminRepository
}

var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesAdminMetricsRepository)(nil)
var _ domain.UserRolesAdminRepository = (*UserRolesAdminMetricsRepository)(nil)

func NewUserRolesAdminMetricsRepository(repo domain.UserRolesAdminRepository) *UserRolesAdminMetricsRepository {
	return &UserRolesAdminMetricsRepository{
		repo:                       repo,
		BaseOwnedMetricsRepository: metrics.NewBaseOwnedMetricsRepository[*domain.Role, string, string]("UserRolesAdminRepository", repo),
	}
}
