package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// UserAdminSaveTraceInteractor implements the usecase.UserAdminSaveUseCase interface,
// serving as a distributed tracing decorator leveraging OpenTelemetry API primitives.
// Automatically spans across business scenario persistence workflows to capture runtime states.
type UserAdminSaveTraceInteractor struct {
	*telemetry.BaseTelemetry                              // Embedded base framework-level telemetry orchestrator
	next                     usecase.UserAdminSaveUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                       // Pre-calculated target tracing operational span name
}

// Compile-time interface compliance verification
var _ usecase.UserAdminSaveUseCase = (*UserAdminSaveTraceInteractor)(nil)

// NewUserAdminSaveTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer.
func NewUserAdminSaveTraceUseCase(ucName string, next usecase.UserAdminSaveUseCase) *UserAdminSaveTraceInteractor {
	return &UserAdminSaveTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Save", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Save executes the underlying domain persistence layer wrapped safely within an isolated OTel child span context.
func (uast *UserAdminSaveTraceInteractor) Save(ctx context.Context, model *domain.User) (*domain.User, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := uast.StartSpan(ctx, uast.spanName)
	defer span.End()

	// Capture contextual object meta payload parameters into tracking attributes safely
	if utils.IsNil(model) {
		span.SetAttributes(attribute.String("param.entity_name", "nil_pointer"))
	} else {
		span.SetAttributes(attribute.String("param.entity_name", model.Name))
	}

	res, err := uast.next.Save(ctx, model)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Save_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
