package domain

import (
	"hash/fnv"
	"strconv"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	libauditdom "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
)

// Role encapsulates the core enterprise domain RBAC aggregate role entity state,
// boundary validation checkpoints, cryptographic hash calculation sequences, and data-auditing serialization blueprints.
type Role struct {
	ID          string
	Name        string
	Description string
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Compile-time interface compliance verifications
var _ libdom.Entity[string] = (*Role)(nil)
var _ libdom.SoftDeleteEntity[bool] = (*Role)(nil)
var _ libauditdom.Auditable = (*Role)(nil)
var _ repository.AuditableEntity[string] = (*Role)(nil)

// NewEmptyRole allocates a baseline empty Role structure blueprint handle.
func NewEmptyRole() *Role {
	return &Role{}
}

// NewRole acts as a complete constructor function orchestrating deterministic value mapping over all structural entity components.
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

// GetID extracts the current unique sequence identification string token handle.
func (r *Role) GetID() string {
	return r.ID
}

// SetID maps an explicit external identifier configuration sequence onto the entity state.
func (r *Role) SetID(id string) {
	r.ID = id
}

// IsExists reports whether the aggregate currently commands a valid structural primary persistence index identity.
func (r *Role) IsExists() bool {
	return r.ID != ""
}

// GetDeleted returns the raw internal boolean marker indicating archived state properties.
func (r *Role) GetDeleted() bool {
	return r.Deleted
}

// SetDeleted modifies the internal structural soft-deletion lifecycle flag state parameters.
func (r *Role) SetDeleted(deleted bool) {
	r.Deleted = deleted
}

// IsDeleted reports a fast semantic flag query evaluating archived identity status metadata.
func (r *Role) IsDeleted() bool {
	return r.Deleted
}

// BeforeCreate triggers automated time-sorted UUIDv7 token injections and updates localized temporary auditing clocks.
func (r *Role) BeforeCreate() error {
	if err := libdom.AssignUUIDv7(r); err != nil {
		return errs.NewBllError("Role.BeforeCreate", "default before create failed", err)
	}

	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now()
	}
	r.UpdatedAt = time.Now()

	return nil
}

// BeforeChange updates internal clock fields prior to writing modifications downstream.
func (r *Role) BeforeChange() error {
	r.UpdatedAt = time.Now()

	return nil
}

// ValidateCreate asserts boundary invariants to guarantee zero identity index drift happens during execution.
func (r *Role) ValidateCreate() error {
	if r.ID != "" {
		return errs.NewBllValidateError("Role.ValidateCreate", "id must be empty", nil)
	}
	if r.Name == "" {
		return errs.NewBllValidateError("Role.ValidateCreate", "name cannot be empty", nil)
	}

	return nil
}

// ValidateChange asserts presence requirements forcing clean non-empty identification criteria fields during state modifications.
func (r *Role) ValidateChange() error {
	if r.ID == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "id cannot be empty", nil)
	}
	if r.Name == "" {
		return errs.NewBllValidateError("Role.ValidateChange", "name cannot be empty", nil)
	}

	return nil
}

// GetInternalTypeName unpacks low-level structural reflection names detailing active models taxonomies.
func (r *Role) GetInternalTypeName() string {
	return utils.GetFullTypeName(r)
}

// GetTypeName outputs the flat canonical string alias representation characterizing this class configuration node.
func (r *Role) GetTypeName() string {
	return "Role"
}

// GetTypeDescription outputs short descriptive contextual labels.
func (r *Role) GetTypeDescription() string {
	return "Role model"
}

// GetInstanceID proxies unique index extractions to fulfill underlying framework diagnostic capabilities contracts.
func (r *Role) GetInstanceID() string {
	return r.ID
}

// GetInstanceName maps generic literal indicators for centralized system monitoring operations.
func (r *Role) GetInstanceName() string {
	return r.Name
}

// HashCode serializes the cumulative scalar state parameters into a stable cryptographic FNV-1a uint32 representation.
func (r *Role) HashCode() uint32 {
	h := fnv.New32a()

	_, _ = h.Write([]byte(r.ID))
	_, _ = h.Write([]byte(r.Name))
	if r.Deleted {
		_, _ = h.Write([]byte{1})
	} else {
		_, _ = h.Write([]byte{0})
	}
	_, _ = h.Write([]byte(r.CreatedAt.Format(time.RFC3339)))
	_, _ = h.Write([]byte(r.UpdatedAt.Format(time.RFC3339)))

	return h.Sum32()
}

// ToAuditMap serializes entity fields into a structured audit registry schema for external log-streaming broker pipelines.
func (r *Role) ToAuditMap() map[string]*libauditdom.AuditField {
	res := make(map[string]*libauditdom.AuditField)

	res["id"] = libauditdom.NewAuditField(r.ID, "УИЭ")
	res["name"] = libauditdom.NewAuditField(r.Name, "Наименование")
	res["description"] = libauditdom.NewAuditField(r.Description, "Описание")
	res["deleted"] = libauditdom.NewAuditField(strconv.FormatBool(r.Deleted), "Признак soft delete")
	res["created_at"] = libauditdom.NewAuditField(r.CreatedAt.Format(time.RFC3339), "Создано")
	res["updated_at"] = libauditdom.NewAuditField(r.UpdatedAt.Format(time.RFC3339), "Изменено")

	return res
}
