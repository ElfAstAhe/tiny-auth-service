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

type RegisterTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.RegisterUseCase
	spanName string
}

var _ usecase.RegisterUseCase = (*RegisterTraceInteractor)(nil)

func NewRegisterTraceUseCase(ucName string, next usecase.RegisterUseCase) *RegisterTraceInteractor {
	return &RegisterTraceInteractor{
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
		spanName:      fmt.Sprintf("%s.Register", ucName),
		next:          next,
	}
}

func (rti *RegisterTraceInteractor) Register(ctx context.Context, username string, password string) (*domain.User, error) {
	ctx, span := rti.StartSpan(ctx, rti.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.username", username))

	res, err := rti.next.Register(ctx, username, password)
	if err != nil {
		span.AddEvent("Register_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
