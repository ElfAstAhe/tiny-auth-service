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

type UserAdminSaveTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.UserAdminSaveUseCase
	spanName string
}

func NewUserAdminSaveTraceUseCase(ucName string, next usecase.UserAdminSaveUseCase) *UserAdminSaveTraceInteractor {
	return &UserAdminSaveTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Save", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (uast *UserAdminSaveTraceInteractor) Save(ctx context.Context, model *domain.User) (*domain.User, error) {
	ctx, span := uast.StartSpan(ctx, uast.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.entity_name", model.Name))

	res, err := uast.next.Save(ctx, model)
	if err != nil {
		span.AddEvent("Save_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
