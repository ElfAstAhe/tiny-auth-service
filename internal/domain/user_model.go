package domain

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	auditdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"golang.org/x/exp/slices"
)

type User struct {
	ID           string
	Name         string
	Type         string
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
var _ auditdomain.Auditable = (*User)(nil)

func NewEmptyUser() *User {
	return &User{
		Roles: make([]*Role, 0),
	}
}

func NewUser(id, name, userType, passwordHash, publicKey, privateKey string, active, deleted bool, createdAt time.Time, roles ...*Role) *User {
	return &User{
		ID:           id,
		Name:         name,
		Type:         userType,
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
	if err := u.validateCommon(); err != nil {
		return errs.NewBllValidateError("User.ValidateCreate", "common validation failed", err)
	}

	return nil
}

func (u *User) ValidateChange() error {
	if u.ID == "" {
		return errs.NewBllValidateError("User.ValidateChange", "id cannot be empty", nil)
	}
	if err := u.validateCommon(); err != nil {
		return errs.NewBllValidateError("User.ValidateChange", "common validation failed", err)
	}

	return nil
}

func (u *User) validateCommon() error {
	if u.Name == "" {
		return errs.NewBllValidateError("User.ValidateChange", "name cannot be empty", nil)
	}
	if err := validateUserType(u.Type); err != nil {
		return errs.NewBllValidateError("User.ValidateChange", "type validation failed", err)
	}
	if u.PasswordHash == "" {
		return errs.NewBllValidateError("User.ValidateChange", "password hash cannot be empty", nil)
	}

	return nil
}

func (u *User) GetInternalTypeName() string {
	return utils.GetFullTypeName(u)
}

func (u *User) GetTypeName() string {
	return "User"
}

func (u *User) GetTypeDescription() string {
	return "User model"
}

func (u *User) GetInstanceID() string {
	return u.ID
}

func (u *User) GetInstanceName() string {
	return u.Name
}

func (u *User) HashCode() uint32 {
	h := fnv.New32a()

	h.Write([]byte(u.ID))
	h.Write([]byte(u.Name))
	h.Write([]byte(u.Type))
	h.Write([]byte(u.PasswordHash))
	h.Write([]byte(u.PublicKey))
	h.Write([]byte(u.PrivateKey))
	h.Write([]byte(strconv.FormatBool(u.Active)))
	h.Write([]byte(strconv.FormatBool(u.Deleted)))
	h.Write([]byte(u.CreatedAt.Format(time.RFC3339)))
	h.Write([]byte(u.UpdatedAt.Format(time.RFC3339)))

	roleIDs := domain.EntitiesToIDList(u.Roles)
	slices.Sort(roleIDs)
	for _, roleID := range roleIDs {
		h.Write([]byte(roleID))
	}

	return h.Sum32()
}

func (u *User) ToAuditMap() map[string]string {
	res := make(map[string]string)

	res["id"] = u.ID
	res["name"] = u.Name
	res["type"] = u.Type
	res["password_hash"] = u.PasswordHash
	res["public_key"] = u.PublicKey
	res["private_key"] = u.PrivateKey
	res["active"] = strconv.FormatBool(u.Active)
	res["deleted"] = strconv.FormatBool(u.Deleted)
	res["created_at"] = u.CreatedAt.Format(time.RFC3339)
	res["updated_at"] = u.UpdatedAt.Format(time.RFC3339)

	roles := make([]string, 0, len(u.Roles))
	for _, role := range u.Roles {
		roles = append(roles, fmt.Sprintf("%s.%s", role.ID, role.Name))
	}
	slices.Sort(roles)
	res["roles"] = strings.Join(roles, ",")

	return res
}
