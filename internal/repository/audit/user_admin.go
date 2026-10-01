package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	libauditdom "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminAuditRepository implements the domain.UserAdminRepository interface,
// acting as a non-invasive database auditing decorator component.
// It intercepts structural administrative data mutations to automatically extract, map, and stream audit records via a broker client.
type UserAdminAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.User, string]                            // Generic framework-level audit base decorator structure
	next                                                      domain.UserAdminRepository // Downstream storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verification
var _ domain.UserAdminRepository = (*UserAdminAuditRepository)(nil)

// NewUserAdminRepository acts as a factory constructor mounting audit logging proxies over administrative user storage adapters.
func NewUserAdminRepository(
	source string,
	next domain.UserAdminRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *UserAdminAuditRepository {
	res := &UserAdminAuditRepository{
		next: next,
	}
	res.BaseAuditCRUDRepository = repository.NewBaseAuditCRUDRepository[*domain.User, string](
		next,
		source,
		res.mapEntityToAuditable,
		auditClient,
		log,
	)

	return res
}

// FindByName proxies name-based query execution straight to the downstream data adapter, bypassing change auditing layers.
func (uaa *UserAdminAuditRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	return uaa.next.FindByName(ctx, name)
}

// mapEntityToAuditable transforms an internal user domain object into a transmittable framework audit contract interface.
func (uaa *UserAdminAuditRepository) mapEntityToAuditable(entity *domain.User) libauditdom.Auditable {
	return entity
}
