package audit

import (
	"context"
	"time"

	"github.com/ElfAstAhe/tiny-audit-service/pkg/api/http/audit/v1/models"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client/rest"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type AuthFacadeImpl struct {
	next        *facade.AuthFacadeImpl
	auditClient rest.AuditClient[*models.AuthAuditDTO]
}

var _ facade.AuthFacade = (*AuthFacadeImpl)(nil)

func NewAuthFacade(
	auditClient rest.AuditClient[*models.AuthAuditDTO],
	next *facade.AuthFacadeImpl,
) *AuthFacadeImpl {
	return &AuthFacadeImpl{
		next:        next,
		auditClient: auditClient,
	}
}

func (aaf *AuthFacadeImpl) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.Login(ctx, login)

	// аудит
	data := aaf.buildAudit(login, res, err)

	// отправка
	_ = aaf.auditClient.Audit(data)

	return res, err
}

func (aaf *AuthFacadeImpl) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.LoginSimple(ctx, login)

	// аудит
	data := aaf.buildAudit(login, res, err)

	// отправка
	_ = aaf.auditClient.Audit(data)

	return res, err
}

func (aaf *AuthFacadeImpl) buildAudit(req *dto.LoginDTO, res *dto.LoggedInDTO, err error) *models.AuthAuditDTO {
	builder := utils.NewAuthAuditBuilder().
		NewInstance().
		WithSource("tiny-auth-service").
		WithEventDate(time.Now()).
		WithEvent(rest.AuthEventLogin).
		WithRequestID("").
		WithTraceID("").
		WithUsername(req.Username)

	if res != nil {
		builder.WithAccessToken(res.Token).
			WithRefreshToken(res.RefreshToken)
	}

	return builder.WithStatus(aaf.toDtoStatus(err)).
		Build()
}

func (aaf *AuthFacadeImpl) toDtoStatus(err error) string {
	switch err == nil {
	case true:
		return rest.AuditStatusSuccess
	default:
		return rest.AuditStatusFail
	}
}
