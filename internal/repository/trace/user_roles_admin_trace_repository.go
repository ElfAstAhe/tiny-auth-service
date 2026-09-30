package trace

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesAdminTraceRepository struct {
	*trace.BaseOwnedTraceRepository[*domain.Role, string, string]
	repo domain.UserRolesAdminRepository
}

var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesAdminTraceRepository)(nil)
var _ domain.UserRolesAdminRepository = (*UserRolesAdminTraceRepository)(nil)

func NewUserRolesAdminTraceRepository(repo domain.UserRolesAdminRepository) *UserRolesAdminTraceRepository {
	return &UserRolesAdminTraceRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: trace.NewBaseOwnedTraceRepository[*domain.Role, string, string]("UserRolesAdminRepository", repo),
	}
}
