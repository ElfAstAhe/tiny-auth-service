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

// UserAdminTraceRepository implements libdomain.CRUDRepository and domain.UserAdminRepository interfaces,
// acting as a non-invasive distributed tracing decorator on top of the concrete administrative user storage adapter.
// It manages OpenTelemetry span lifecycles for high-throughput administrative lookup operations.
type UserAdminTraceRepository struct {
	*trace.BaseCRUDTraceRepository[*domain.User, string]                            // Generic framework-level trace base decorator structure
	repo                                                 domain.UserAdminRepository // Downstream core administrative database persistence adapter layer
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.User, string] = (*UserAdminTraceRepository)(nil)
var _ domain.UserAdminRepository = (*UserAdminTraceRepository)(nil)

// NewUserAdminTraceRepository acts as a factory constructor mounting tracing middleware primitives over an administrative user data adapter.
func NewUserAdminTraceRepository(repo domain.UserAdminRepository) *UserAdminTraceRepository {
	return &UserAdminTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: trace.NewBaseCRUDTraceRepository[*domain.User, string]("UserAdminRepository", repo),
	}
}

// FindByName intercepts administrative lookup sequences by username parameters to execute data retrieval wrapped inside an OTel trace span boundary.
func (uat *UserAdminTraceRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	ctx, span := uat.StartSpan(ctx, fmt.Sprintf("%s.FindByName", uat.GetRepositoryName()))
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely (fully aligned with the param. prefix)
	span.SetAttributes(attribute.String("param.name", name))

	res, err := uat.repo.FindByName(ctx, name)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
