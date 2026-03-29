package telemetry

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type LoginSimpleTraceInteractor struct {
	*telemetry.BaseTelemetry
	next     usecase.LoginSimpleUseCase
	spanName string
}

func NewLoginSimpleTraceUseCase(ucName string, next usecase.LoginUseCase) *LoginTraceInteractor {
	return &LoginTraceInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.LoginSimple", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

func (lt *LoginTraceInteractor) LoginSimple(ctx context.Context, username string, encryptedPassword string) (*jwt.Token, *jwt.Token, error) {
	ctx, span := lt.StartSpan(ctx, lt.spanName)
	defer span.End()

	span.SetAttributes(attribute.String("param.username", username))

	token, refreshToken, err := lt.next.Login(ctx, username, encryptedPassword)
	if err != nil {
		span.AddEvent("Login_simple_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return token, refreshToken, err
}
