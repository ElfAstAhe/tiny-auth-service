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

// RegisterTraceInteractor implements the usecase.RegisterUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts user registration workflows to track identity creation lifecycles and safely append metadata context.
type RegisterTraceInteractor struct {
	*telemetry.BaseTelemetry                         // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.RegisterUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                  // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.RegisterUseCase = (*RegisterTraceInteractor)(nil)

// NewRegisterTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for identity registration operations.
func NewRegisterTraceUseCase(ucName string, next usecase.RegisterUseCase) *RegisterTraceInteractor {
	return &RegisterTraceInteractor{
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
		spanName:      fmt.Sprintf("%s.Register", ucName),
		next:          next,
	}
}

// Register executes a traced wrapper sequence over the underlying core business scenario "Register" account creation procedure.
func (rti *RegisterTraceInteractor) Register(ctx context.Context, username string, password string) (*domain.User, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := rti.StartSpan(ctx, rti.spanName)
	defer span.End()

	// Capture contextual non-sensitive telemetry parameters into tracking attributes safely (password parameters are explicitly omitted)
	span.SetAttributes(attribute.String("param.username", username))

	res, err := rti.next.Register(ctx, username, password)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Register_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
