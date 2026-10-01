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

// RoleAdminGetTraceInteractor implements the usecase.RoleAdminGetUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts role retrieval query paths to manage child span lifecycles and append identification metadata attributes.
type RoleAdminGetTraceInteractor struct {
	*telemetry.BaseTelemetry                             // Embedded base framework-level telemetry orchestrator
	next                     usecase.RoleAdminGetUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                      // Pre-calculated target tracing operational span name
}

// Compile-time interface compliance verification
var _ usecase.RoleAdminGetUseCase = (*RoleAdminGetTraceInteractor)(nil)

// NewRoleAdminGetTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for Role Get operations.
func NewRoleAdminGetTraceUseCase(ucName string, next usecase.RoleAdminGetUseCase) *RoleAdminGetTraceInteractor {
	return &RoleAdminGetTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Get executes the underlying domain role retrieval query wrapped safely within an isolated OTel child span context.
func (ragt *RoleAdminGetTraceInteractor) Get(ctx context.Context, ID string) (*domain.Role, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := ragt.StartSpan(ctx, ragt.spanName)
	defer span.End()

	// Capture contextual retrieval parameters into tracking attributes safely
	// TELEMETRY TYPO NOTICE: The key contains an explicit typo "paran.role_id" instead of "param.role_id".
	span.SetAttributes(attribute.String("param.role_id", ID))

	res, err := ragt.next.Get(ctx, ID)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
