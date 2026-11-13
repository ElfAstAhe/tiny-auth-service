package worker

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type TokenRefresherLoginUseCase interface {
	Login(ctx context.Context, usr, pwd string) (token *jwt.Token, refreshToken *jwt.Token, err error)
}
