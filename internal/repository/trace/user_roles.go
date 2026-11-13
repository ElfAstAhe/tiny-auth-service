package trace

import (
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserRolesRepository implements libdomain.OwnedRepository and domain.UserRolesRepository interfaces,
// acting as a non-invasive distributed tracing decorator component for owned role entity storage operations.
// It encapsulates OpenTelemetry span lifecycles, injecting tracing boundaries across repository aggregate execution vectors.
type UserRolesRepository struct {
	*trace.BaseOwnedTraceRepository[*domain.Role, string, string]                            // Generic framework-level trace base decorator structure for owned aggregates
	repo                                                          domain.UserRolesRepository // Downstream data storage persistence or metrics layer execution implementation path
}

// Compile-time interface compliance verifications
var _ libdomain.OwnedRepository[*domain.Role, string, string] = (*UserRolesRepository)(nil)
var _ domain.UserRolesRepository = (*UserRolesRepository)(nil)

// NewUserRolesRepository acts as a factory constructor mounting OpenTelemetry distributed tracing proxies over user-roles storage adapters.
func NewUserRolesRepository(repo domain.UserRolesRepository) *UserRolesRepository {
	return &UserRolesRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: trace.NewBaseOwnedTraceRepository[*domain.Role, string, string]("UserRolesRepository", repo),
	}
}
