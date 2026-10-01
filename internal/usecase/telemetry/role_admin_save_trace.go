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

// RoleAdminSaveTraceInteractor implements the usecase.RoleAdminSaveUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts role entity persistence workflows to record operational trace spans and capture metadata state attributes.
type RoleAdminSaveTraceInteractor struct {
	*telemetry.BaseTelemetry                              // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.RoleAdminSaveUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                       // Pre-calculated target tracing operational span name
}

// Compile-time interface compliance verification
var _ usecase.RoleAdminSaveUseCase = (*RoleAdminSaveTraceInteractor)(nil)

// NewRoleAdminSaveTraceUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for Role Save operations.
func NewRoleAdminSaveTraceUseCase(ucName string, next usecase.RoleAdminSaveUseCase) *RoleAdminSaveTraceInteractor {
	return &RoleAdminSaveTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Save", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Save executes the underlying domain role persistence layer wrapped safely within an isolated OTel child span context.
func (ast *RoleAdminSaveTraceInteractor) Save(ctx context.Context, model *domain.Role) (*domain.Role, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := ast.StartSpan(ctx, ast.spanName)
	defer span.End()

	// Capture contextual object meta payload parameters into tracking attributes safely
	if utils.IsNil(model) {
		span.SetAttributes(attribute.String("param.model_name", "nil_pointer"))
	} else {
		span.SetAttributes(attribute.String("param.model_name", model.Name))
	}

	res, err := ast.next.Save(ctx, model)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Save_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
