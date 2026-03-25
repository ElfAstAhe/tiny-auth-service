package dto

import (
	"time"
)

// RoleDTO represent role model
// @Description Role DTO
type RoleDTO struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
} // @name RoleDTO
