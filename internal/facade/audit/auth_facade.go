package audit

import (
	"context"
	"time"

	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditdto "github.com/ElfAstAhe/tiny-audit-service/pkg/client/dto"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type AuthFacadeImpl struct {
	next        *facade.AuthFacadeImpl
	source      string
	auditClient client.AuthAuditClient
}

var _ facade.AuthFacade = (*AuthFacadeImpl)(nil)

func NewAuthFacade(
	auditClient client.AuthAuditClient,
	source string,
	next *facade.AuthFacadeImpl,
) *AuthFacadeImpl {
	return &AuthFacadeImpl{
		source:      source,
		next:        next,
		auditClient: auditClient,
	}
}

func (aaf *AuthFacadeImpl) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.Login(ctx, login)

	// аудит
	data := aaf.buildAudit(ctx, login, res, err)

	// отправка
	_ = aaf.auditClient.Audit(data)

	return res, err
}

func (aaf *AuthFacadeImpl) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.LoginSimple(ctx, login)

	// аудит
	data := aaf.buildAudit(ctx, login, res, err)

	// отправка
	_ = aaf.auditClient.Audit(data)

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
