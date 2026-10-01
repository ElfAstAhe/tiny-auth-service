package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

// UserAdminRepository defines the administrative domain persistence contract managing CRUD operations
// and specialized index query routines targeting core User aggregate entities with elevated privileges.
type UserAdminRepository interface {
	domain.CRUDRepository[*User, string] // Core library generic blueprint interface for CRUD persistence operations

	// FindByName executes a dedicated index-based administrative lookup sequence utilizing a unique string name criterion.
	FindByName(context.Context, string) (*User, error)
}
