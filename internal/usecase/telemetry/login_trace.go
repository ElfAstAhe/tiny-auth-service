package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// LoginTraceInteractor implements the usecase.LoginUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts user authentication flows to track session creation lifecycles and safely append metadata context.
type LoginTraceInteractor struct {
	*telemetry.BaseTelemetry                      // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.LoginUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string               // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.LoginUseCase = (*LoginTraceInteractor)(nil)

// NewLoginTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for identity login operations.
func NewLoginTraceUseCase(ucName string, next usecase.LoginUseCase) *LoginTraceInteractor {
	return &LoginTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Login", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Login executes a traced wrapper sequence over the underlying core business scenario "Login" session initialization procedure.
func (lt *LoginTraceInteractor) Login(ctx context.Context, username string, encryptedPassword string) (*jwt.Token, *jwt.Token, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := lt.StartSpan(ctx, lt.spanName)
	defer span.End()

	// Capture contextual non-sensitive telemetry parameters into tracking attributes safely (crypto password payloads are explicitly omitted)
	span.SetAttributes(attribute.String("param.username", username))

	token, refreshToken, err := lt.next.Login(ctx, username, encryptedPassword)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Login_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return token, refreshToken, err
}
