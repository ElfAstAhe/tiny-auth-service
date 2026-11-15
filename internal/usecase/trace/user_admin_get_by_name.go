package trace

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// UserAdminGetByNameInteractor implements the usecase.UserAdminGetByNameUseCase interface,
// serving as a non-invasive distributed tracing decorator leveraging OpenTelemetry API.
// It intercepts user identity query flows to track runtime lifecycle state using name-based lookup semantics.
type UserAdminGetByNameInteractor struct {
	*telemetry.BaseTelemetry                                   // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.UserAdminGetByNameUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                            // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.UserAdminGetByNameUseCase = (*UserAdminGetByNameInteractor)(nil)

// NewUserAdminGetByNameUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for name-based lookup operations.
func NewUserAdminGetByNameUseCase(ucName string, next usecase.UserAdminGetByNameUseCase) *UserAdminGetByNameInteractor {
	return &UserAdminGetByNameInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Get executes a traced wrapper sequence over the underlying core business scenario "Get" retrieval procedure using string keys.
func (gnt *UserAdminGetByNameInteractor) Get(ctx context.Context, name string) (*domain.User, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := gnt.StartSpan(ctx, gnt.spanName)
	defer span.End()

	// Capture high-cardinality string lookup attributes into tracking span metadata safely
	span.SetAttributes(attribute.String("param.name", name))

	res, err := gnt.next.Get(ctx, name)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
