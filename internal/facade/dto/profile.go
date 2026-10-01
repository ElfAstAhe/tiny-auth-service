package dto

import (
	"time"
)

// ProfileDTO encapsulates the structural data transfer object profile blueprint payload
// returned to consumers requesting validated user identity information records.
type ProfileDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	PublicKey string    `json:"public_key"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Roles     []string  `json:"roles"`
} // @name ProfileDTO
