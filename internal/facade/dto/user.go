package dto

import (
	"time"
)

type UserDTO struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"password_hash,omitempty"`
	PublicKey    string    `json:"public_key,omitempty"`
	PrivateKey   string    `json:"private_key,omitempty"`
	Active       bool      `json:"active"`
	Deleted      bool      `json:"deleted"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	Roles []*RoleDTO `json:"roles,omitempty"`
}
