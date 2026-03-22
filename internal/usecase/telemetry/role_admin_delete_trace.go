package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type RoleAdminDeleteTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.RoleAdminDeleteUseCase
	spanName string
}

func NewRoleAdminDeleteTraceUseCase(ucName string, next usecase.RoleAdminDeleteUseCase) *RoleAdminDeleteTraceInteractor {
	return &RoleAdminDeleteTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Delete", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (radt *RoleAdminDeleteTraceInteractor) Delete(ctx context.Context, ID string) error {
	ctx, span := radt.StartSpan(ctx, radt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.role_id", ID))

	err := radt.next.Delete(ctx, ID)
	if err != nil {
		span.AddEvent("Delete_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
