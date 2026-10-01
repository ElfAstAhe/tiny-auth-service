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

// RoleAdminListTraceInteractor implements the usecase.RoleAdminListUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts high-throughput query lists for roles to inject tracking spans and capture metadata attributes.
type RoleAdminListTraceInteractor struct {
	*telemetry.BaseTelemetry                              // Embedded base framework-level telemetry orchestrator
	next                     usecase.RoleAdminListUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                       // Pre-calculated target tracing operational span name
}

// Compile-time interface compliance verification
var _ usecase.RoleAdminListUseCase = (*RoleAdminListTraceInteractor)(nil)

// NewRoleAdminListTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for Role List operations.
func NewRoleAdminListTraceUseCase(ucName string, next usecase.RoleAdminListUseCase) *RoleAdminListTraceInteractor {
	return &RoleAdminListTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.List", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// List executes the underlying domain role collection retrieval wrapped safely within an isolated OTel child span context.
func (alt *RoleAdminListTraceInteractor) List(ctx context.Context, limit, offset int) ([]*domain.Role, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := alt.StartSpan(ctx, alt.spanName)
	defer span.End()

	// Capture contextual pagination parameters into tracking attributes safely
	span.SetAttributes(
		attribute.Int("param.limit", limit),
		attribute.Int("param.offset", offset),
	)

	res, err := alt.next.List(ctx, limit, offset)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("List_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
