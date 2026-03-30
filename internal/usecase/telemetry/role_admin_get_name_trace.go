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

type RoleAdminGetNameTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.RoleAdminGetNameUseCase
	spanName string
}

var _ usecase.RoleAdminGetNameUseCase = (*RoleAdminGetNameTraceInteractor)(nil)

func NewROleAdminGetNameTraceUseCase(ucName string, next usecase.RoleAdminGetNameUseCase) *RoleAdminGetNameTraceInteractor {
	return &RoleAdminGetNameTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (gnt *RoleAdminGetNameTraceInteractor) Get(ctx context.Context, name string) (*domain.Role, error) {
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
