package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// ChangePasswordTraceInteractor implements the usecase.ChangePasswordUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts password mutation workflows to track account security lifecycles and record failures.
type ChangePasswordTraceInteractor struct {
	*telemetry.BaseTelemetry                               // Embedded framework-level telemetry engine core orchestrator
	spanName                 string                        // Pre-formatted OpenTelemetry specification span name
	next                     usecase.ChangePasswordUseCase // The encapsulated downstream active core business usecase logic
}

// Compile-time interface compliance verification
var _ usecase.ChangePasswordUseCase = (*ChangePasswordTraceInteractor)(nil)

// NewChangePasswordTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for password modification operations.
func NewChangePasswordTraceUseCase(ucName string, next usecase.ChangePasswordUseCase) *ChangePasswordTraceInteractor {
	return &ChangePasswordTraceInteractor{
		next:          next,
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
		spanName:      fmt.Sprintf("%s.ChangePassword", ucName),
	}
}

// ChangePassword executes a traced wrapper sequence over the underlying core business scenario "ChangePassword" account state modification.
func (cpt *ChangePasswordTraceInteractor) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := cpt.StartSpan(ctx, cpt.spanName)
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely (fully aligned with the param. prefix)
	span.SetAttributes(attribute.String("param.user_id", userID))

	err := cpt.next.ChangePassword(ctx, userID, oldPassword, newPassword)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("ChangePassword_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
