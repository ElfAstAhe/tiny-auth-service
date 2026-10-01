package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// ChangeKeysTraceInteractor implements the usecase.ChangeKeysUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts cryptographic key rotation workflows to track identity security cycles and record failure states.
type ChangeKeysTraceInteractor struct {
	*telemetry.BaseTelemetry                           // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.ChangeKeysUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                    // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.ChangeKeysUseCase = (*ChangeKeysTraceInteractor)(nil)

// NewChangeKeysTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for key modification operations.
func NewChangeKeysTraceUseCase(ucName string, next usecase.ChangeKeysUseCase) *ChangeKeysTraceInteractor {
	return &ChangeKeysTraceInteractor{
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
		next:          next,
		spanName:      fmt.Sprintf("%s.ChangeKeys", ucName),
	}
}

// ChangeKeys executes a traced wrapper sequence over the underlying core business scenario "ChangeKeys" cryptography state modification.
func (ckt *ChangeKeysTraceInteractor) ChangeKeys(ctx context.Context, userID string) (string, string, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := ckt.StartSpan(ctx, ckt.spanName)
	defer span.End()

	// Capture contextual parameters into tracking attributes safely (returned public/private keys are explicitly omitted)
	// TELEMETRY ALIGNMENT NOTICE: The key currently utilizes "user.id". Consider refactoring to "param.user_id" to maintain unified project-wide semantic conventions.
	span.SetAttributes(attribute.String("user.id", userID))

	publicKey, privateKey, err := ckt.next.ChangeKeys(ctx, userID)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("ChangeKeys_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return publicKey, privateKey, err
}
