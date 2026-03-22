package repository

import (
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserRolesAdminTraceRepository struct {
	*repository.BaseOwnedTraceRepository[*domain.Role, string, string]
	repo domain.UserRolesAdminRepository
}

func NewUserRolesAdminTraceRepository(repo domain.UserRolesAdminRepository) *UserRolesAdminTraceRepository {
	return &UserRolesAdminTraceRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: repository.NewBaseOwnedTraceRepository[*domain.Role, string, string]("UserRolesAdminRepository", repo),
	}
}
