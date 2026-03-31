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

type RoleAdminGetTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.RoleAdminGetUseCase
	spanName string
}

var _ usecase.RoleAdminGetUseCase = (*RoleAdminGetTraceInteractor)(nil)

func NewRoleAdminGetTraceUseCase(ucName string, next usecase.RoleAdminGetUseCase) *RoleAdminGetTraceInteractor {
	return &RoleAdminGetTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (ragt *RoleAdminGetTraceInteractor) Get(ctx context.Context, ID string) (*domain.Role, error) {
	ctx, span := ragt.StartSpan(ctx, ragt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("paran.role_id", ID))

	res, err := ragt.next.Get(ctx, ID)
	if err != nil {
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
