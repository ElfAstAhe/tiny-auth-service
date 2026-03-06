package domain

import (
	"time"

	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/google/uuid"
)

type Role struct {
	ID          string
	Name        string
	Description string
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewEmptyRole() *Role {
	return &Role{}
}

func NewRole(id string, name string, description string, deleted bool, createdAt time.Time, updatedAt time.Time) *Role {
	return &Role{
		ID:          id,
		Name:        name,
		Description: description,
		Deleted:     deleted,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (r *Role) GetID() string {
	return r.ID
}

func (r *Role) SetID(id string) {
	r.ID = id
}

func (r *Role) IsExists() bool {
	return r.ID != ""
}

func (r *Role) GetDeleted() bool {
	return r.Deleted
}

func (r *Role) SetDeleted(deleted bool) {
	r.Deleted = deleted
}

func (r *Role) IsDeleted() bool {
	return r.Deleted
}

func (r *Role) BeforeCreate() error {
	newID, err := uuid.NewRandom()
	if err != nil {
		return errs.NewBllError("Role.BeforeCreate", "generate new id", err)
	}

	r.ID = newID.String()
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now()
	}
	r.UpdatedAt = time.Now()

	return nil
}

func (r *Role) BeforeChange() error {
	r.UpdatedAt = time.Now()

	return nil
}

func (r *Role) ValidateCreate() error {
	if r.ID != "" {
		return errs.NewBllValidateError("Role.ValidateCreate", "id must be empty", nil)
	}
	if r.Name == "" {
		return errs.NewBllValidateError("Role.ValidateCreate", "name cannot be empty", nil)
	}

	return nil
}

func (r *Role) ValidateChange() error {
	if r.ID == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "id cannot be empty", nil)
	}
	if r.Name == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "name cannot be empty", nil)
	}

	return nil
}
