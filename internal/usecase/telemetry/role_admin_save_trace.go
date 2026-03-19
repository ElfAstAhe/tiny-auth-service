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

type RoleAdminSaveTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.RoleAdminSaveUseCase
	spanName string
}

func NewRoleAdminSaveTraceUseCase(ucName string, next usecase.RoleAdminSaveUseCase) *RoleAdminSaveTraceInteractor {
	return &RoleAdminSaveTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Save", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (ast *RoleAdminSaveTraceInteractor) Save(ctx context.Context, model *domain.Role) (*domain.Role, error) {
	ctx, span := ast.StartSpan(ctx, ast.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.model_name", model.Name))

	res, err := ast.next.Save(ctx, model)
	if err != nil {
		span.AddEvent("Save_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
