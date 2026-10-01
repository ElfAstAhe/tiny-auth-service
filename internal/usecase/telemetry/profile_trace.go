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

// ProfileTraceInteractor implements the usecase.ProfileUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts user profile requests to track metrics and inject contextual logging attributes.
type ProfileTraceInteractor struct {
	*telemetry.BaseTelemetry                        // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.ProfileUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                 // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.ProfileUseCase = (*ProfileTraceInteractor)(nil)

// NewProfileTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for Profile operations.
func NewProfileTraceUseCase(ucName string, next usecase.ProfileUseCase) *ProfileTraceInteractor {
	return &ProfileTraceInteractor{
		spanName:      fmt.Sprintf("%s.Get", ucName),
		next:          next,
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Get executes a traced wrapper sequence over the underlying core business scenario "Get" profile retrieval procedure.
func (pt *ProfileTraceInteractor) Get(ctx context.Context, username string) (*domain.User, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := pt.StartSpan(ctx, pt.spanName)
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely
	span.SetAttributes(attribute.String("param.username", username))

	res, err := pt.next.Get(ctx, username)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
