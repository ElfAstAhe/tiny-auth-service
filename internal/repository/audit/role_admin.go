package audit

import (
	"context"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	libextdom "github.com/ElfAstAhe/go-service-template/pkg/ext/domain"
	libauderrs "github.com/ElfAstAhe/tiny-audit-service/pkg/errs"
	libaudrepo "github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	libaudutl "github.com/ElfAstAhe/tiny-audit-service/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminRepository implements the domain.RoleAdminRepository interface,
// acting as a non-invasive database auditing decorator component for administrative roles.
// It intercepts data mutations to automatically extract, map, and stream audit records via a broker client.
type RoleAdminRepository struct {
	*libaudrepo.AuditCRUD[*domain.Role, *libextdom.DataAuditOutbox, string]
	next domain.RoleAdminRepository // Downstream storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verification
var _ libdom.CRUDRepository[*domain.Role, string] = (*RoleAdminRepository)(nil)
var _ domain.RoleAdminRepository = (*RoleAdminRepository)(nil)

// NewRoleAdminRepository acts as a factory constructor mounting audit logging proxies over administrative role storage adapters.
func NewRoleAdminRepository(
	source string,
	next domain.RoleAdminRepository,
	outboxRepo libextdom.OutboxRepository[*libextdom.DataAuditOutbox, string],
) (*RoleAdminRepository, error) {
	// instance
	res := &RoleAdminRepository{
		next: next,
	}
	// embed
	embed, err := libaudrepo.NewAuditCRUD[*domain.Role, *libextdom.DataAuditOutbox, string](
		libaudrepo.WithAuditSource[*domain.Role, *libextdom.DataAuditOutbox, string](source),
		libaudrepo.WithAuditNext[*domain.Role, *libextdom.DataAuditOutbox, string](next),
		libaudrepo.WithAuditOutboxBuilder[*domain.Role, *libextdom.DataAuditOutbox, string](libaudutl.DefaultBuildDataAuditOutbox),
		libaudrepo.WithAuditOutboxRepository[*domain.Role, *libextdom.DataAuditOutbox, string](outboxRepo),
	)
	if err != nil {
		return nil, libauderrs.NewAuditError("NewRoleRepository", "unable to create audit CRUD embed", err)
	}
	// setup
	res.AuditCRUD = embed

	return res, nil
}

// FindByName proxies name-based query execution straight to the downstream data adapter, bypassing change auditing layers.
func (raa *RoleAdminRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	return raa.next.FindByName(ctx, name)
}
