package audit

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditdto "github.com/ElfAstAhe/tiny-audit-service/pkg/client/dto"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// Deprecated: AuthFacadeImpl оставлен как образец использования audit rest client
type AuthFacadeImpl struct {
	next        facade.AuthFacade
	source      string
	auditClient client.AuthAuditClient
	logger      logger.Logger
}

var _ facade.AuthFacade = (*AuthFacadeImpl)(nil)

func NewAuthFacadeRest(
	auditClient client.AuthAuditClient,
	source string,
	next facade.AuthFacade,
	log logger.Logger,
) *AuthFacadeImpl {
	return &AuthFacadeImpl{
		source:      source,
		next:        next,
		auditClient: auditClient,
		logger:      log.GetLogger("AUTH_FACADE"),
	}
}

func (aaf *AuthFacadeImpl) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.Login(ctx, login)

	// аудит
	data := aaf.buildAudit(ctx, login, res, err)

	// отправка
	err = aaf.auditClient.Audit(data)
	if err != nil {
		aaf.logger.Errorf("error audit: %v", err)
	}

	return res, err
}

func (aaf *AuthFacadeImpl) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.LoginSimple(ctx, login)

	// аудит
	data := aaf.buildAudit(ctx, login, res, err)

	// отправка
	err = aaf.auditClient.Audit(data)
	if err != nil {
		aaf.logger.Errorf("error audit: %v", err)
	}

	return res, err
}

func (aaf *AuthFacadeImpl) buildAudit(ctx context.Context, req *dto.LoginDTO, res *dto.LoggedInDTO, err error) *auditdto.AuthAuditDTO {
	// common
	builder := utils.NewAuthAuditBuilder().
		WithSource(aaf.source).
		WithEventDate(time.Now()).
		WithEvent(auditdto.AuthEventLogin).
		WithUsername(req.Username).
		// request
		WithRequestID(utils.RequestIDFromContext(ctx)).
		WithTraceID(utils.TraceIDFromContext(ctx))

	// tokens
	if res != nil {
		builder.WithAccessToken(res.Token).
			WithRefreshToken(res.RefreshToken)
	}

	// result
	return builder.WithStatus(aaf.toAuditStatus(err)).
		Build()
}

func (aaf *AuthFacadeImpl) toAuditStatus(err error) string {
	switch err == nil {
	case true:
		return auditdto.AuditStatusSuccess
	default:
		return auditdto.AuditStatusFail
	}
}
