package domain

import (
	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type UserRolesAdminRepository interface {
	domain.OwnedRepository[*Role, string, string]
}
