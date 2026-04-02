package trace

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesTraceRepository struct {
	*repository.BaseOwnedTraceRepository[*domain.Role, string, string]
	repo domain.UserRolesRepository
}

var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesTraceRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesTraceRepository)(nil)

func NewUserRolesTraceRepository(repo domain.UserRolesRepository) *UserRolesTraceRepository {
	return &UserRolesTraceRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: repository.NewBaseOwnedTraceRepository[*domain.Role, string, string]("UserRolesRepository", repo),
	}
}
