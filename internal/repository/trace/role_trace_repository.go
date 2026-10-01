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

// RoleTraceRepository implements libdomain.CRUDRepository and domain.RoleRepository interfaces,
// acting as a non-invasive distributed tracing decorator on top of the concrete role storage adapter.
// It manages OpenTelemetry span lifecycles for high-throughput role name-based lookups.
type RoleTraceRepository struct {
	*trace.BaseCRUDTraceRepository[*domain.Role, string]                       // Generic framework-level trace base decorator structure
	repo                                                 domain.RoleRepository // Downstream core role database persistence adapter layer
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleTraceRepository)(nil)
var _ domain.RoleRepository = (*RoleTraceRepository)(nil)

// NewRoleTraceRepository acts as a factory constructor mounting tracing middleware primitives over a core role data adapter.
func NewRoleTraceRepository(repo domain.RoleRepository) *RoleTraceRepository {
	return &RoleTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: trace.NewBaseCRUDTraceRepository[*domain.Role, string]("RoleRepository", repo),
	}
}

// FindByName intercepts lookup sequences by role name parameters to execute data retrieval wrapped inside an OTel trace span boundary.
func (rtr *RoleTraceRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	ctx, span := rtr.StartSpan(ctx, fmt.Sprintf("%s.FindByName", rtr.GetRepositoryName()))
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely (fully aligned with the param. prefix)
	span.SetAttributes(attribute.String("param.name", name))

	res, err := rtr.repo.FindByName(ctx, name)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
