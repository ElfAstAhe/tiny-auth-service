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

type UserAdminGetTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.UserAdminGetUseCase
	spanName string
}

func NewUserAdminGetTraceUseCase(ucName string, next usecase.UserAdminGetNameUseCase) *UserAdminGetTraceInteractor {
	return &UserAdminGetTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (agt *UserAdminGetTraceInteractor) Get(ctx context.Context, ID string) (*domain.User, error) {
	ctx, span := agt.StartSpan(ctx, agt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.user_id", ID))

	res, err := agt.next.Get(ctx, ID)
	if err != nil {
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
