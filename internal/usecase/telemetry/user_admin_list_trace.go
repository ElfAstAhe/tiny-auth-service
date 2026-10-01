package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// UserAdminListTraceInteractor implements the usecase.UserAdminListUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts high-throughput query lists to inject tracking spans and capture metadata attributes.
type UserAdminListTraceInteractor struct {
	*telemetry.BaseTelemetry                              // Embedded base framework-level telemetry orchestrator
	next                     usecase.UserAdminListUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                       // Pre-calculated target tracing operational span name
}

// Compile-time interface compliance verification
var _ usecase.UserAdminListUseCase = (*UserAdminListTraceInteractor)(nil)

// NewUserAdminListTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for List operations.
func NewUserAdminListTraceUseCase(ucName string, next usecase.UserAdminListUseCase) *UserAdminListTraceInteractor {
	return &UserAdminListTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.List", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// List executes the underlying domain collection retrieval wrapped safely within an isolated OTel child span context.
func (ualt *UserAdminListTraceInteractor) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := ualt.StartSpan(ctx, ualt.spanName)
	defer span.End()

	// Capture contextual pagination parameters into tracking attributes safely
	span.SetAttributes(attribute.Int("param.limit", limit), attribute.Int("param.offset", offset))

	res, err := ualt.next.List(ctx, limit, offset)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("List_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
