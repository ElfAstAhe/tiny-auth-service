package usecase

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type LoginTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     LoginUseCase
	spanName string
}

func NewLoginTraceUseCase(ucName string, next LoginUseCase) *LoginTraceInteractor {
	return &LoginTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Login", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (lt *LoginTraceInteractor) Login(ctx context.Context, username string, encryptedPassword string) (*jwt.Token, *jwt.Token, error) {
	ctx, span := lt.StartSpan(ctx, lt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.username", username))

	token, refreshToken, err := lt.next.Login(ctx, username, encryptedPassword)
	if err != nil {
		span.AddEvent("Login_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return token, refreshToken, err
}
