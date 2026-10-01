package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAuditRepository implements the domain.UserRepository interface,
// acting as a non-invasive database auditing decorator component for user entities.
// It intercepts destructive data mutations to automatically extract, map, and stream audit records via a broker client.
type UserAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.User, string]                       // Generic framework-level audit base decorator structure
	next                                                      domain.UserRepository // Downstream storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verification
var _ domain.UserRepository = (*UserAuditRepository)(nil)

// NewUserRepository acts as a factory constructor mounting audit logging proxies over standard user storage adapters.
func NewUserRepository(
	source string,
	next domain.UserRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *UserAuditRepository {
	res := &UserAuditRepository{
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
func (uar *UserAuditRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	return uar.next.FindByName(ctx, name)
}

// mapEntityToAuditable transforms an internal user domain object into a transmittable framework audit contract interface.
func (uar *UserAuditRepository) mapEntityToAuditable(entity *domain.User) auditlibdomain.Auditable {
	return entity
}
