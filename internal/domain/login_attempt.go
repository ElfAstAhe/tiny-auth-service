package domain

import (
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

// LoginAttempt represent table login_attempts
type LoginAttempt struct {
	ID          string
	ContentType string
	Data        string
	Sent        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Compile-time interface compliance verifications
var _ libdom.Entity[string] = (*LoginAttempt)(nil)

// NewEmptyLoginAttempt allocates a baseline empty LoginAttempt structure blueprint handle.
func NewEmptyLoginAttempt() *LoginAttempt {
	return &LoginAttempt{}
}

// NewLoginAttempt acts as a complete constructor function orchestrating deterministic value mapping over all structural entity components.
func NewLoginAttempt(id string, contentType string, data string, createdAt time.Time, sent bool) *LoginAttempt {
	return &LoginAttempt{
		ID:          id,
		ContentType: contentType,
		Data:        data,
		CreatedAt:   createdAt,
		Sent:        sent,
	}
}

// GetID extracts the current unique sequence identification string token handle.
func (la *LoginAttempt) GetID() string {
	return la.ID
}

// SetID maps an explicit external identifier configuration sequence onto the entity state.
func (la *LoginAttempt) SetID(id string) {
	la.ID = id
}

// IsExists reports whether the aggregate currently commands a valid structural primary persistence index identity.
func (la *LoginAttempt) IsExists() bool {
	return la.ID != ""
}

// BeforeCreate triggers automated time-sorted UUIDv7 token injections and updates localized temporary auditing clocks.
func (la *LoginAttempt) BeforeCreate() error {
	if err := libdom.AssignUUIDv7(la); err != nil {
		return errs.NewBllError("LoginAttempt.BeforeCreate", "default before create failed", err)
	}

	if la.CreatedAt.IsZero() {
		la.CreatedAt = time.Now()
	}
	la.UpdatedAt = time.Now()

	return nil
}

// BeforeChange updates internal clock fields prior to writing modifications downstream.
func (la *LoginAttempt) BeforeChange() error {
	la.UpdatedAt = time.Now()

	return nil
}

// ValidateCreate asserts boundary invariants to guarantee zero identity index drift happens during execution.
func (la *LoginAttempt) ValidateCreate() error {
	if la.ID != "" {
		return errs.NewBllValidateError("LoginAttempt.ValidateCreate", "id must be empty", nil)
	}
	if la.ContentType == "" {
		return errs.NewBllValidateError("LoginAttempt.ValidateCreate", "contentType required", nil)
	}
	if la.Data == "" {
		return errs.NewBllValidateError("LoginAttempt.ValidateCreate", "data required", nil)
	}

	return nil
}

// ValidateChange asserts presence requirements forcing clean non-empty identification criteria fields during state modifications.
func (la *LoginAttempt) ValidateChange() error {
	if la.ID == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "id cannot be empty", nil)
	}
	if la.ContentType == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "contentType required", nil)
	}
	if la.Data == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "data required", nil)
	}

	return nil
}
