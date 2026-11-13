package trace

import (
	"context"
	"fmt"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// RoleAdminRepository implements libdomain.CRUDRepository and domain.RoleAdminRepository interfaces,
// acting as a non-invasive distributed tracing decorator on top of the concrete administrative role storage adapter.
// It manages OpenTelemetry span lifecycles for high-throughput administrative lookup operations.
type RoleAdminRepository struct {
	*trace.CRUD[*domain.Role, string]                            // Generic framework-level trace base decorator structure
	repo                              domain.RoleAdminRepository // Downstream core administrative database persistence adapter layer
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleAdminRepository)(nil)
var _ domain.RoleAdminRepository = (*RoleAdminRepository)(nil)

// NewRoleAdminRepository acts as a factory constructor mounting tracing middleware primitives over an administrative role data adapter.
func NewRoleAdminRepository(repo domain.RoleAdminRepository) *RoleAdminRepository {
	return &RoleAdminRepository{
		repo: repo,
		CRUD: trace.NewCRUD[*domain.Role, string]("RoleAdminRepository", repo),
	}
}

// FindByName intercepts administrative lookup sequences by role name parameters to execute data retrieval wrapped inside an OTel trace span boundary.
//
//goland:noinspection DuplicatedCode
func (rat *RoleAdminRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	ctx, span := rat.StartSpan(ctx, fmt.Sprintf("%s.FindByName", rat.GetRepositoryName()))
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely (fully aligned with the param. prefix)
	span.SetAttributes(attribute.String("param.name", name))

	res, err := rat.repo.FindByName(ctx, name)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
