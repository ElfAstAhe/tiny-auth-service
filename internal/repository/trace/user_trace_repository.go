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

// UserTraceRepository implements libdomain.CRUDRepository and domain.UserRepository interfaces,
// acting as a non-invasive distributed tracing decorator on top of the concrete user storage adapter.
// Encapsulates OpenTelemetry span lifecycles for high-throughput repository name-based lookups.
type UserTraceRepository struct {
	*trace.BaseCRUDTraceRepository[*domain.User, string]                       // Generic framework-level trace base decorator structure
	repo                                                 domain.UserRepository // Downstream core database persistence adapter layer implementation
}

// Compile-time interface compliance verifications
var _ libdomain.CRUDRepository[*domain.User, string] = (*UserTraceRepository)(nil)
var _ domain.UserRepository = (*UserTraceRepository)(nil)

// NewUserTraceRepository acts as a factory constructor mounting tracing middleware primitives over a core user data adapter.
func NewUserTraceRepository(repo domain.UserRepository) *UserTraceRepository {
	return &UserTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: trace.NewBaseCRUDTraceRepository[*domain.User, string]("UserRepository", repo),
	}
}

// FindByName intercepts lookup sequences by username parameters to execute data retrieval wrapped inside an OTel trace span boundary.
//
//goland:noinspection DuplicatedCode
func (utr *UserTraceRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	ctx, span := utr.StartSpan(ctx, fmt.Sprintf("%s.FindByName", utr.GetRepositoryName()))
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely (fully aligned with the param. prefix)
	span.SetAttributes(attribute.String("param.name", name))

	res, err := utr.repo.FindByName(ctx, name)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
