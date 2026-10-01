package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

// UserRepository defines the domain persistence contract managing CRUD operations
// and specialized query routines targeting core User aggregate entities.
type UserRepository interface {
	domain.CRUDRepository[*User, string] // Core library generic blueprint interface for CRUD persistence operations

	// FindByName executes a dedicated index-based retrieval lookup sequence utilizing a unique string login criterion.
	FindByName(ctx context.Context, login string) (*User, error)
}
