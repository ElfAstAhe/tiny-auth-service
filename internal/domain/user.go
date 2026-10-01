package domain

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	auditdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"golang.org/x/exp/slices"
)

// User encapsulates the core enterprise domain identity aggregate state, mappings, invariants verification checkpoints,
// cryptographic hash-sum calculation sequences, and data-auditing serialization blueprints.
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

// Compile-time interface compliance verifications
var _ libdomain.Entity[string] = (*User)(nil)
var _ libdomain.SoftDeleteEntity[bool] = (*User)(nil)
var _ auditdomain.Auditable = (*User)(nil)
var _ repository.AuditableEntity[string] = (*User)(nil)

// NewEmptyUser allocates a baseline User structure stub initializing empty child nested slice structures.
func NewEmptyUser() *User {
	return &User{
		Roles: make([]*Role, 0),
	}
}

// NewUser acts as a complete constructor function orchestrating deterministic value mapping over all structural entity components.
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

// GetID extracts the current unique sequence identification string token handle.
func (u *User) GetID() string {
	return u.ID
}

// SetID maps an explicit external identifier configuration sequence onto the entity state.
func (u *User) SetID(id string) {
	u.ID = id
}

// IsExists reports whether the aggregate currently commands a valid structural primary persistence index identity.
func (u *User) IsExists() bool {
	return u.ID != ""
}

// GetDeleted returns the raw internal boolean marker indicating archived state properties.
func (u *User) GetDeleted() bool {
	return u.Deleted
}

// SetDeleted modifies the internal structural soft-deletion lifecycle flag state parameters.
func (u *User) SetDeleted(deleted bool) {
	u.Deleted = deleted
}

// IsDeleted reports a fast semantic flag query evaluating archived identity status metadata.
func (u *User) IsDeleted() bool {
	return u.Deleted
}

// BeforeCreate triggers automated time-sorted UUIDv7 token injections and updates localized temporary auditing clocks.
func (u *User) BeforeCreate() error {
	if err := libdomain.AssignUUIDv7(u); err != nil {
		return errs.NewBllError("User.BeforeCreate", "default before create failed", err)
	}

	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	return nil
}

// BeforeChange updates internal clock fields prior to writing modifications downstream.
func (u *User) BeforeChange() error {
	u.UpdatedAt = time.Now()

	return nil
}

// ValidateCreate asserts boundary invariants to guarantee zero identity index drift happens during execution.
func (u *User) ValidateCreate() error {
	if u.ID != "" {
		return errs.NewBllValidateError("User.ValidateCreate", "id must be empty", nil)
	}
	if err := u.validateCommon(); err != nil {
		return errs.NewBllValidateError("User.ValidateCreate", "common validation failed", err)
	}

	return nil
}

// ValidateChange asserts presence requirements forcing clean non-empty identification criteria fields during state modifications.
func (u *User) ValidateChange() error {
	if u.ID == "" {
		return errs.NewBllValidateError("User.ValidateChange", "id cannot be empty", nil)
	}
	if err := u.validateCommon(); err != nil {
		return errs.NewBllValidateError("User.ValidateChange", "common validation failed", err)
	}

	return nil
}

// validateCommon reviews mutual cross-cutting presence constraints across base literal components.
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

// GetInternalTypeName unpacks low-level structural reflection names detailing active models taxonomies.
func (u *User) GetInternalTypeName() string {
	return utils.GetFullTypeName(u)
}

// GetTypeName outputs the flat canonical string alias representation characterizing this class configuration node.
func (u *User) GetTypeName() string {
	return "User"
}

// GetTypeDescription outputs short descriptive contextual labels.
func (u *User) GetTypeDescription() string {
	return "User model"
}

// GetInstanceID proxies unique index extractions to fulfill underlying framework diagnostic capabilities contracts.
func (u *User) GetInstanceID() string {
	return u.ID
}

// GetInstanceName maps generic literal indicators for centralized system monitoring operations.
func (u *User) GetInstanceName() string {
	return u.Name
}

// HashCode serializes the cumulative scalar state parameters into a stable cryptographic FNV-1a uint32 representation.
func (u *User) HashCode() uint32 {
	h := fnv.New32a()

	_, _ = h.Write([]byte(u.ID))
	_, _ = h.Write([]byte(u.Name))
	_, _ = h.Write([]byte(u.Type))
	_, _ = h.Write([]byte(u.PasswordHash))
	_, _ = h.Write([]byte(u.PublicKey))
	_, _ = h.Write([]byte(u.PrivateKey))
	_, _ = h.Write([]byte(strconv.FormatBool(u.Active)))
	_, _ = h.Write([]byte(strconv.FormatBool(u.Deleted)))
	_, _ = h.Write([]byte(u.CreatedAt.Format(time.RFC3339)))
	_, _ = h.Write([]byte(u.UpdatedAt.Format(time.RFC3339)))

	roleIDs := libdomain.EntitiesToIDList(u.Roles)
	//nolint:govet // broken inline analyzer
	slices.Sort(roleIDs)
	for _, roleID := range roleIDs {
		_, _ = h.Write([]byte(roleID))
	}

	return h.Sum32()
}

// ToAuditMap serializes entity fields into a structured audit registry schema for external log-streaming broker pipelines.
func (u *User) ToAuditMap() map[string]*auditdomain.AuditField {
	res := make(map[string]*auditdomain.AuditField)

	res["id"] = auditdomain.NewAuditField(u.ID, "УИЭ")
	res["name"] = auditdomain.NewAuditField(u.Name, "Наименование")
	res["type"] = auditdomain.NewAuditField(u.Type, "Тип")
	res["password_hash"] = auditdomain.NewAuditField(u.PasswordHash, "hash пароля")
	res["public_key"] = auditdomain.NewAuditField(u.PublicKey, "RSA публичный ключ")
	res["private_key"] = auditdomain.NewAuditField(u.PrivateKey, "RSA скрытый ключ")
	res["active"] = auditdomain.NewAuditField(strconv.FormatBool(u.Active), "Признак пользователь активирован")
	res["deleted"] = auditdomain.NewAuditField(strconv.FormatBool(u.Deleted), "Признак soft delete")
	res["created_at"] = auditdomain.NewAuditField(u.CreatedAt.Format(time.RFC3339), "Создано")
	res["updated_at"] = auditdomain.NewAuditField(u.UpdatedAt.Format(time.RFC3339), "Изменено")

	roles := make([]string, 0, len(u.Roles))
	for _, role := range u.Roles {
		roles = append(roles, fmt.Sprintf("%s.%s", role.ID, role.Name))
	}
	//nolint:govet // broken inline analyzer
	slices.Sort(roles)
	res["roles"] = auditdomain.NewAuditField(strings.Join(roles, ","), "Роли")

	return res
}
