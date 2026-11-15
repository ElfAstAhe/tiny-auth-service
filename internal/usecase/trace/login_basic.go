package trace

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// LoginBasicInteractor implements the usecase.LoginBasicUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts simplified credentials authentication flows to manage child span lifecycles and safely track telemetry metadata.
type LoginBasicInteractor struct {
	*telemetry.BaseTelemetry                           // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.LoginBasicUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                    // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.LoginBasicUseCase = (*LoginBasicInteractor)(nil)

// NewLoginBasicUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for simplified login operations.
func NewLoginBasicUseCase(ucName string, next usecase.LoginBasicUseCase) *LoginBasicInteractor {
	return &LoginBasicInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Login", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Login executes a traced wrapper sequence over the underlying core business scenario "Login" simplified session initialization procedure.
func (lbt *LoginBasicInteractor) Login(ctx context.Context, username string, encryptedPassword string) (*jwt.Token, *jwt.Token, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := lbt.StartSpan(ctx, lbt.spanName)
	defer span.End()

	// Capture contextual non-sensitive telemetry parameters into tracking attributes safely (crypto password payloads are explicitly omitted)
	span.SetAttributes(attribute.String("param.username", username))

	token, refreshToken, err := lbt.next.Login(ctx, username, encryptedPassword)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Login_simple_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return token, refreshToken, err
}
