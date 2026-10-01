package trace

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRolesTraceRepository implements libdomain.OwnedRepository and domain.UserRolesRepository interfaces,
// acting as a non-invasive distributed tracing decorator component for owned role entity storage operations.
// It encapsulates OpenTelemetry span lifecycles, injecting tracing boundaries across repository aggregate execution vectors.
type UserRolesTraceRepository struct {
	*trace.BaseOwnedTraceRepository[*domain.Role, string, string]                            // Generic framework-level trace base decorator structure for owned aggregates
	repo                                                          domain.UserRolesRepository // Downstream data storage persistence or metrics layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesTraceRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesTraceRepository)(nil)

// NewUserRolesTraceRepository acts as a factory constructor mounting OpenTelemetry distributed tracing proxies over user-roles storage adapters.
func NewUserRolesTraceRepository(repo domain.UserRolesRepository) *UserRolesTraceRepository {
	return &UserRolesTraceRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: trace.NewBaseOwnedTraceRepository[*domain.Role, string, string]("UserRolesRepository", repo),
	}
}
