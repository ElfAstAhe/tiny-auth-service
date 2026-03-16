package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type ChangePasswordTraceInteractor struct {
	*telemetry.BaseTelemetry
	spanName string
	next     ChangePasswordUseCase
}

func NewChangePasswordTraceUseCase(ucName string, next ChangePasswordUseCase) *ChangePasswordTraceInteractor {
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
