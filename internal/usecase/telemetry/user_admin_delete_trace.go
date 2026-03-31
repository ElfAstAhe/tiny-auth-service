package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type UserAdminDeleteTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.UserAdminDeleteUseCase
	spanName string
}

var _ usecase.UserAdminDeleteUseCase = (*UserAdminDeleteTraceInteractor)(nil)

func NewUserAdminDeleteTraceUseCase(ucName string, next usecase.UserAdminDeleteUseCase) *UserAdminDeleteTraceInteractor {
	return &UserAdminDeleteTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Delete", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (adt *UserAdminDeleteTraceInteractor) Delete(ctx context.Context, ID string) error {
	ctx, span := adt.StartSpan(ctx, adt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.user_id", ID))

	err := adt.next.Delete(ctx, ID)
	if err != nil {
		span.AddEvent("Delete_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
