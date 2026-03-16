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

type RoleAdminListTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.RoleAdminListUseCase
	spanName string
}

func NewRoleAdminListTraceUseCase(ucName string, next usecase.RoleAdminListUseCase) *RoleAdminListTraceInteractor {
	return &RoleAdminListTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.List", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (alt *RoleAdminListTraceInteractor) List(ctx context.Context, limit, offset int) ([]*domain.Role, error) {
	ctx, span := alt.StartSpan(ctx, alt.spanName)
	defer span.End()

	span.SetAttributes(attribute.Int("limit", limit), attribute.Int("offset", offset))

	res, err := alt.next.List(ctx, limit, offset)
	if err != nil {
		span.AddEvent("List_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
