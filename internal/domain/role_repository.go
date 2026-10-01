package domain

import (
	"context"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
)

// RoleRepository defines the domain persistence contract managing CRUD operations
// and specialized query routines targeting core Role aggregate entities.
type RoleRepository interface {
	libdom.CRUDRepository[*Role, string] // Core library generic blueprint interface for CRUD persistence operations

	// FindByName executes a dedicated index-based retrieval lookup sequence utilizing a unique string role name criterion.
	FindByName(context.Context, string) (*Role, error)
}
