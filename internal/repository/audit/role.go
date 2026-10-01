package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAuditRepository implements the domain.RoleRepository interface,
// acting as a non-invasive database auditing decorator component for standard user roles.
// It intercepts data mutations to automatically extract, map, and stream audit records via a broker client.
type RoleAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.Role, string]                       // Generic framework-level audit base decorator structure
	next                                                      domain.RoleRepository // Downstream storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verification
var _ domain.RoleRepository = (*RoleAuditRepository)(nil)

// NewRoleRepository acts as a factory constructor mounting audit logging proxies over standard role storage adapters.
func NewRoleRepository(
	source string,
	next domain.RoleRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *RoleAuditRepository {
	res := &RoleAuditRepository{
		next: next,
	}
	res.BaseAuditCRUDRepository = repository.NewBaseAuditCRUDRepository[*domain.Role, string](
		next,
		source,
		res.mapEntityToAuditable,
		auditClient,
		log,
	)

	return res
}

// FindByName proxies name-based query execution straight to the downstream data adapter, bypassing change auditing layers.
func (rar *RoleAuditRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	return rar.next.FindByName(ctx, name)
}

// mapEntityToAuditable transforms an internal role domain object into a transmittable framework audit contract interface.
func (rar *RoleAuditRepository) mapEntityToAuditable(entity *domain.Role) auditlibdomain.Auditable {
	return entity
}
