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

type ProfileTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.ProfileUseCase
	spanName string
}

func NewProfileTraceUseCase(ucName string, next usecase.ProfileUseCase) *ProfileTraceInteractor {
	return &ProfileTraceInteractor{
		spanName:      fmt.Sprintf("%s.Get", ucName),
		next:          next,
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (pt *ProfileTraceInteractor) Get(ctx context.Context, username string) (*domain.User, error) {
	ctx, span := pt.StartSpan(ctx, pt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.username", username))

	res, err := pt.next.Get(ctx, username)
	if err != nil {
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
