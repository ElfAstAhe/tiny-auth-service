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

type UserAdminGetNameTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.UserAdminGetNameUseCase
	spanName string
}

func NewUserAdminGetNameTraceUseCase(ucName string, next usecase.UserAdminGetNameUseCase) *UserAdminGetNameTraceInteractor {
	return &UserAdminGetNameTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (gnt *UserAdminGetNameTraceInteractor) Get(ctx context.Context, name string) (*domain.User, error) {
	ctx, span := gnt.StartSpan(ctx, gnt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.name", name))

	res, err := gnt.next.Get(ctx, name)
	if err != nil {
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
