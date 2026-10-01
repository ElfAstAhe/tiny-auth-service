package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// UserAdminDeleteTraceInteractor implements the usecase.UserAdminDeleteUseCase interface,
// serving as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts destructive identity eviction flows to track operational runtime states and record failure event logs.
type UserAdminDeleteTraceInteractor struct {
	*telemetry.BaseTelemetry                                // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.UserAdminDeleteUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                         // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.UserAdminDeleteUseCase = (*UserAdminDeleteTraceInteractor)(nil)

// NewUserAdminDeleteTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for Delete operations.
func NewUserAdminDeleteTraceUseCase(ucName string, next usecase.UserAdminDeleteUseCase) *UserAdminDeleteTraceInteractor {
	return &UserAdminDeleteTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Delete", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Delete executes a traced wrapper sequence over the underlying core business scenario "Delete" removal procedure.
func (adt *UserAdminDeleteTraceInteractor) Delete(ctx context.Context, ID string) error {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := adt.StartSpan(ctx, adt.spanName)
	defer span.End()

	// Capture high-cardinality payload identifiers into tracking span metadata safely
	span.SetAttributes(attribute.String("param.user_id", ID))

	err := adt.next.Delete(ctx, ID)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Delete_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
