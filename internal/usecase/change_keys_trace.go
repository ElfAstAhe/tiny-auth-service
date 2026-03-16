package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type ChangeKeysTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     ChangeKeysUseCase
	spanName string
}

func NewChangeKeysTraceInteractor(ucName string, next ChangeKeysUseCase) *ChangeKeysTraceInteractor {
	return &ChangeKeysTraceInteractor{
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
		next:          next,
		spanName:      fmt.Sprintf("%s.ChangeKeys", ucName),
	}
}

func (ckt *ChangeKeysTraceInteractor) ChangeKeys(ctx context.Context, userID string) (string, string, error) {
	ctx, span := ckt.StartSpan(ctx, ckt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("user.id", userID))

	publicKey, privateKey, err := ckt.next.ChangeKeys(ctx, userID)
	if err != nil {
		span.AddEvent("ChangeKeys_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return publicKey, privateKey, err
}
