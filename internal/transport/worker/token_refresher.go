package worker

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/auth"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/worker"
)

type TokenRefresher struct {
	*worker.BaseTokenRefresher
	jwtHelper     *helper.JWTHelper
	loginSimpleUC usecase.LoginSimpleUseCase
	creds         *config.ServiceCredentialsConfig
}

var _ auth.TokenProvider = (*TokenRefresher)(nil)
var _ libworker.Scheduler = (*TokenRefresher)(nil)

func NewTokenRefresher(
	jwtHelper *helper.JWTHelper,
	simpleLoginUC usecase.LoginSimpleUseCase,
	creds *config.ServiceCredentialsConfig,
	conf *worker.BaseTokenRefresherConfig,
	log logger.Logger,
) *TokenRefresher {
	res := &TokenRefresher{
		jwtHelper:     jwtHelper,
		loginSimpleUC: simpleLoginUC,
		creds:         creds,
	}

	res.BaseTokenRefresher = worker.NewBaseTokenRefresher(
		conf,
		res.tokenRefreshAction,
		log,
	)

	return res
}

func (tr *TokenRefresher) tokenRefreshAction(ctx context.Context, eventTime time.Time) (string, error) {
	tr.GetLogger().Debugf("token refresher timer event %s start", eventTime.Format(time.DateTime))
	defer tr.GetLogger().Debugf("token refresher timer event %s finish", eventTime.Format(time.DateTime))

	token, _, err := tr.loginSimpleUC.Login(ctx, tr.creds.Username, tr.creds.Password)
	if err != nil {
		return "", err
	}

	tokenStr, err := tr.jwtHelper.BuildTokenStr(token)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
