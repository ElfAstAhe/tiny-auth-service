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

type UserAdminListTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.UserAdminListUseCase
	spanName string
}

func NewUserAdminListTraceUseCase(ucName string, next usecase.UserAdminListUseCase) *UserAdminListTraceInteractor {
	return &UserAdminListTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.List", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (ualt *UserAdminListTraceInteractor) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	ctx, span := ualt.StartSpan(ctx, ualt.spanName)
	defer span.End()

	span.SetAttributes(attribute.Int("param.limit", limit), attribute.Int("param.offset", offset))

	res, err := ualt.next.List(ctx, limit, offset)
	if err != nil {
		span.AddEvent("List_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
