package domain

import (
	"hash/fnv"
	"strconv"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	auditdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type Role struct {
	ID          string
	Name        string
	Description string
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var _ domain.Entity[string] = (*Role)(nil)
var _ domain.SoftDeleteEntity[bool] = (*Role)(nil)
var _ auditdomain.Auditable = (*Role)(nil)

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
	if err := defaultBeforeCreate(r); err != nil {
		return errs.NewBllError("Role.BeforeCreate", "default before create failed", err)
	}

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

func (r *Role) GetInternalTypeName() string {
	return utils.GetFullTypeName(r)
}

func (r *Role) GetTypeName() string {
	return "Role"
}

func (r *Role) GetTypeDescription() string {
	return "Role model"
}

func (r *Role) GetInstanceID() string {
	return r.ID
}

func (r *Role) GetInstanceName() string {
	return r.Name
}

func (r *Role) HashCode() uint32 {
	h := fnv.New32a()

	h.Write([]byte(r.ID))
	h.Write([]byte(r.Name))
	if r.Deleted {
		h.Write([]byte{1})
	} else {
		h.Write([]byte{0})
	}
	h.Write([]byte(r.CreatedAt.Format(time.RFC3339)))
	h.Write([]byte(r.UpdatedAt.Format(time.RFC3339)))

	return h.Sum32()
}

func (r *Role) ToAuditMap() map[string]string {
	res := make(map[string]string)

	res["id"] = r.ID
	res["name"] = r.Name
	res["description"] = r.Description
	res["deleted"] = strconv.FormatBool(r.Deleted)
	res["created_at"] = r.CreatedAt.Format(time.RFC3339)
	res["updated_at"] = r.UpdatedAt.Format(time.RFC3339)

	return res
}
