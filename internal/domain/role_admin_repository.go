package domain

import (
	"context"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
)

// RoleAdminRepository defines the administrative domain persistence contract managing CRUD operations
// and specialized index query routines targeting core Role aggregate entities with elevated privileges.
type RoleAdminRepository interface {
	libdom.CRUDRepository[*Role, string] // Core library generic blueprint interface for CRUD persistence operations

	// FindByName executes a dedicated index-based administrative lookup sequence utilizing a unique string key criterion.
	FindByName(ctx context.Context, login string) (*Role, error)
}
