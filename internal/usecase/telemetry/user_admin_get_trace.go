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

// UserAdminGetTraceInteractor implements the usecase.UserAdminGetUseCase interface,
// acting as a non-invasive distributed tracing decorator on top of OpenTelemetry (OTel).
// It intercepts business usecase flows to automatically manage span lifecycles and inject telemetry baggage.
type UserAdminGetTraceInteractor struct {
	*telemetry.BaseTelemetry                             // Framework-level telemetry engine core
	next                     usecase.UserAdminGetUseCase // The wrapped underlying concrete business scenario execution path
	spanName                 string                      // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.UserAdminGetUseCase = (*UserAdminGetTraceInteractor)(nil)

// NewUserAdminGetTraceUseCase acts as a factory constructor mounting a middleware tracing layer on top of a business interactor.
func NewUserAdminGetTraceUseCase(ucName string, next usecase.UserAdminGetUseCase) *UserAdminGetTraceInteractor {
	return &UserAdminGetTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Get executes a traced wrapper sequence over the underlying core business scenario "Get" retrieval procedure.
// Safely tracks span lifecycle states, records exception events, and maps errors into formal infrastructure statuses.
func (agt *UserAdminGetTraceInteractor) Get(ctx context.Context, ID string) (*domain.User, error) {
	// Spawns a child OpenTelemetry span bound seamlessly to the propagation sequence context
	ctx, span := agt.StartSpan(ctx, agt.spanName)
	defer span.End()

	// Capture high-cardinality parameter identifiers into safe span metadata attributes
	span.SetAttributes(attribute.String("param.user_id", ID))

	res, err := agt.next.Get(ctx, ID)
	if err != nil {
		// Log structured system degradation context details directly into the tracing payload
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
