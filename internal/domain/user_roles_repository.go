package domain

import (
	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type UserRolesRepository interface {
	domain.OwnedRepository[*Role, string, string]
}
