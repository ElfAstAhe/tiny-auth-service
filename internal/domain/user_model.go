package domain

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type User struct {
	ID           string
	Name         string
	PasswordHash string
	PublicKey    string
	PrivateKey   string
	Active       bool
	Deleted      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Roles []*Role
}

var _ domain.Entity[string] = (*User)(nil)
var _ domain.SoftDeleteEntity[bool] = (*User)(nil)

func NewEmptyUser() *User {
	return &User{
		Roles: make([]*Role, 0),
	}
}

func NewUser(id, name, passwordHash, publicKey, privateKey string, active, deleted bool, createdAt time.Time, roles ...*Role) *User {
	return &User{
		ID:           id,
		Name:         name,
		PasswordHash: passwordHash,
		PublicKey:    publicKey,
		PrivateKey:   privateKey,
		Active:       active,
		Deleted:      deleted,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
		Roles:        roles,
	}
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) SetID(id string) {
	u.ID = id
}

func (u *User) IsExists() bool {
	return u.ID != ""
}

func (u *User) GetDeleted() bool {
	return u.Deleted
}

func (u *User) SetDeleted(deleted bool) {
	u.Deleted = deleted
}

func (u *User) IsDeleted() bool {
	return u.Deleted
}

func (u *User) BeforeCreate() error {
	if err := defaultBeforeCreate(u); err != nil {
		return errs.NewBllError("User.BeforeCreate", "default before create failed", err)
	}

	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) BeforeChange() error {
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) ValidateCreate() error {
	if u.ID != "" {
		return errs.NewBllValidateError("User.ValidateCreate", "id must be empty", nil)
	}
	if u.Name == "" {
		return errs.NewBllValidateError("User.ValidateCreate", "name cannot be empty", nil)
	}
	if u.PasswordHash == "" {
		return errs.NewBllValidateError("User.ValidateCreate", "password hash cannot be empty", nil)
	}

	return nil
}

func (u *User) ValidateChange() error {
	if u.ID == "" {
		return errs.NewBllValidateError("User.ValidateChange", "id cannot be empty", nil)
	}
	if u.Name == "" {
		return errs.NewBllValidateError("User.ValidateChange", "name cannot be empty", nil)
	}
	if u.PasswordHash == "" {
		return errs.NewBllValidateError("User.ValidateChange", "password hash cannot be empty", nil)
	}

	return nil
}
