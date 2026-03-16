package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type ChangePasswordTraceInteractor struct {
	*telemetry.BaseTelemetry
	spanName string
	next     usecase.ChangePasswordUseCase
}

func NewChangePasswordTraceUseCase(ucName string, next usecase.ChangePasswordUseCase) *ChangePasswordTraceInteractor {
	return &ChangePasswordTraceInteractor{
		next:          next,
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
		spanName:      fmt.Sprintf("%s.ChangePassword", ucName),
	}
}

func (cpt *ChangePasswordTraceInteractor) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	ctx, span := cpt.StartSpan(ctx, cpt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("user.id", userID))

	err := cpt.next.ChangePassword(ctx, userID, oldPassword, newPassword)
	if err != nil {
		span.AddEvent("ChangePassword_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
