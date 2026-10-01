package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RoleAdminAuditRepository implements the domain.RoleAdminRepository interface,
// acting as a non-invasive database auditing decorator component for administrative roles.
// It intercepts data mutations to automatically extract, map, and stream audit records via a broker client.
type RoleAdminAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.Role, string]                            // Generic framework-level audit base decorator structure
	next                                                      domain.RoleAdminRepository // Downstream storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verification
var _ domain.RoleAdminRepository = (*RoleAdminAuditRepository)(nil)

// NewRoleAdminRepository acts as a factory constructor mounting audit logging proxies over administrative role storage adapters.
func NewRoleAdminRepository(
	source string,
	next domain.RoleAdminRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *RoleAdminAuditRepository {
	res := &RoleAdminAuditRepository{
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
func (raa *RoleAdminAuditRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	return raa.next.FindByName(ctx, name)
}

// mapEntityToAuditable transforms an internal role domain object into a transmittable framework audit contract interface.
func (raa *RoleAdminAuditRepository) mapEntityToAuditable(entity *domain.Role) auditlibdomain.Auditable {
	return entity
}
