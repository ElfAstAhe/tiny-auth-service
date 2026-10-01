package trace

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRolesAdminTraceRepository implements libdomain.OwnedRepository and domain.UserRolesAdminRepository interfaces,
// acting as a non-invasive distributed tracing decorator component for administrative owned role aggregate storage operations.
// It encapsulates OpenTelemetry span lifecycles, injecting tracing boundaries across administrative roles management execution vectors.
type UserRolesAdminTraceRepository struct {
	*trace.BaseOwnedTraceRepository[*domain.Role, string, string]                                 // Generic framework-level trace base decorator structure for administrative owned aggregates
	repo                                                          domain.UserRolesAdminRepository // Downstream administrative data storage persistence or metrics layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesAdminTraceRepository)(nil)
var _ domain.UserRolesAdminRepository = (*UserRolesAdminTraceRepository)(nil)

// NewUserRolesAdminTraceRepository acts as a factory constructor mounting OpenTelemetry distributed tracing proxies over administrative user-roles storage adapters.
func NewUserRolesAdminTraceRepository(repo domain.UserRolesAdminRepository) *UserRolesAdminTraceRepository {
	return &UserRolesAdminTraceRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: trace.NewBaseOwnedTraceRepository[*domain.Role, string, string]("UserRolesAdminRepository", repo),
	}
}
