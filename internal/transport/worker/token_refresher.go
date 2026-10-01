package worker

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/auth"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/worker"
)

// TokenRefresher extends the platform's worker.BaseTokenRefresher, implementing
// auth.TokenProvider and libworker.Scheduler contracts to handle active service credentials rotation.
type TokenRefresher struct {
	*worker.BaseTokenRefresher                                  // Embedded platform core base token refresher composition
	jwtHelper                  *helper.JWTHelper                // Framework utility managing serialization of token cryptographic layouts
	loginSimpleUC              usecase.LoginSimpleUseCase       // Simplified business usecase executing identity credentials evaluation
	creds                      *config.ServiceCredentialsConfig // Configuration credentials passport carrying system username and passwords
}

// Compile-time interface compliance verifications
var _ auth.TokenProvider = (*TokenRefresher)(nil)
var _ libworker.Scheduler = (*TokenRefresher)(nil)

// NewTokenRefresher acts as a factory constructor orchestrating full background token rotation worker initialization bounds.
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

// tokenRefreshAction encapsulates execution routines kicked off periodically by ticker events to renew current token assets.
func (tr *TokenRefresher) tokenRefreshAction(ctx context.Context, eventTime time.Time) (string, error) {
	tr.GetLogger().Debugf("token refresher timer event %s start", eventTime.Format(time.DateTime))
	defer tr.GetLogger().Debugf("token refresher timer event %s finish", eventTime.Format(time.DateTime))

	if utils.IsNil(tr.loginSimpleUC) {
		return "", errs.NewCommonError("login use case not provided", nil)
	}

	token, _, err := tr.loginSimpleUC.Login(ctx, tr.creds.Username, tr.creds.Password)
	if err != nil {
		return "", err
	}

	tokenStr, err := tr.jwtHelper.BuildTokenStr(token)
	if err != nil {
		return "", err
	}

	tr.GetLogger().Debugf("token refresher timer event %s got new access token [%s]", eventTime.Format(time.DateTime), tokenStr)

	return tokenStr, nil
}

// SetSimpleLoginUC allows runtime lazy specification adjustments mapping or replacing active business interactor instances.
func (tr *TokenRefresher) SetSimpleLoginUC(useCase usecase.LoginSimpleUseCase) {
	tr.loginSimpleUC = useCase
}
