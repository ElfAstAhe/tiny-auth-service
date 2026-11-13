package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/auth"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/worker"
)

const tokenRefresherNameTemplate = "token-refresher-%s"

// TokenRefresher extends the platform's worker.BaseTokenRefresher, implementing
// auth.TokenProvider and libworker.Scheduler contracts to handle active service credentials rotation.
type TokenRefresher struct {
	*worker.BaseTokenRefresher // Embedded platform core base token refresher composition
	name                       string
	jwtHelper                  *helper.JWTHelper          // Framework utility managing serialization of token cryptographic layouts
	loginUC                    TokenRefresherLoginUseCase // Simplified business usecase executing identity credentials evaluation
	username                   string
	password                   string
	log                        logger.Logger
}

// Compile-time interface compliance verifications
var _ auth.TokenProvider = (*TokenRefresher)(nil)
var _ libworker.Scheduler = (*TokenRefresher)(nil)
var _ container.Runner = (*TokenRefresher)(nil)

// NewTokenRefresher acts as a factory constructor orchestrating full background token rotation worker initialization bounds.
func NewTokenRefresher(options ...TokenRefresherOption) (*TokenRefresher, error) {
	// options
	opts := NewTokenRefreshOptions()
	for _, option := range options {
		option(opts)
	}
	if err := opts.Validate(); err != nil {
		return nil, errs.NewTlCommonError("NewTokenRefresher", "options validation failed", err)
	}
	// instance
	res := &TokenRefresher{
		name:      fmt.Sprintf(tokenRefresherNameTemplate, opts.Name),
		log:       opts.Logger.GetLogger(fmt.Sprintf(tokenRefresherNameTemplate, opts.Name)),
		jwtHelper: opts.JWTHelper,
		loginUC:   opts.LoginUC,
		username:  opts.Username,
		password:  opts.Password,
	}
	// embed
	embed, err := worker.NewBaseTokenRefresher(
		append(
			opts.BaseOptions(),
			worker.WithTokenRefresherTokenRefreshAction(res.tokenRefreshAction),
		)...,
	)
	if err != nil {
		return nil, errs.NewTlCommonError("NewTokenRefresher", "error creating base token refresher instance", err)
	}
	// setup
	res.BaseTokenRefresher = embed

	return res, nil
}

// tokenRefreshAction encapsulates execution routines kicked off periodically by ticker events to renew current token assets.
func (tr *TokenRefresher) tokenRefreshAction(ctx context.Context, eventTime time.Time) (string, error) {
	tr.log.Debugf("token refresher timer event %s start", eventTime.Format(time.DateTime))
	defer tr.log.Debugf("token refresher timer event %s finish", eventTime.Format(time.DateTime))

	if utils.IsNil(tr.loginUC) {
		return "", errs.NewCommonError("login use case not provided", nil)
	}

	token, _, err := tr.loginUC.Login(ctx, tr.username, tr.password)
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
func (tr *TokenRefresher) SetSimpleLoginUC(useCase TokenRefresherLoginUseCase) {
	tr.loginUC = useCase
}
