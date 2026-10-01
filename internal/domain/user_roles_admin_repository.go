package domain

import (
	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

// UserRolesAdminRepository structures the administrative domain contract for handling many-to-many
// owned relationship allocations linking individual users to their respective roles with elevated modification privileges.
type UserRolesAdminRepository interface {
	domain.OwnedRepository[*Role, string, string] // Core library generic blueprint interface for owned child assets operations
}
