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

// AuthFacadeImpl extends the core facade.AuthFacade, acting as a non-invasive synchronous REST-based auditing decorator.
// It intercepts authentication requests, extracts boundary execution metadata, and forwards event logs to the external audit microservice.
//
// Deprecated: AuthFacadeImpl оставлен как образец использования audit rest client
type AuthFacadeImpl struct {
	next        facade.AuthFacade      // Downstream concrete authentication facade implementation handling core session logic
	source      string                 // Context configuration metric defining the identifier string for the origin microservice node
	auditClient client.AuthAuditClient // REST-compatible client manager pushing marshaled audit schemas to targets
	logger      logger.Logger          // Dedicated structured logging handle managing local tracking records
}

// Compile-time interface compliance verification
var _ facade.AuthFacade = (*AuthFacadeImpl)(nil)

// NewAuthFacadeRest acts as a factory constructor mounting required synchronous HTTP auditing decorators over an execution chain.
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

// Login triggers the core session scenario, extracts execution properties, and synchronously pushes authorization statistics into storage logs.
func (aaf *AuthFacadeImpl) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// вызов
	res, err := aaf.next.Login(ctx, login)

	// аудит
	data := aaf.buildAudit(ctx, login, res, err)

	// отправка
	// AUDIT FAULT NOTICE: Network transport degradations inside auditClient do not block downstream user sessions execution.
	err = aaf.auditClient.Audit(data)
	if err != nil {
		aaf.logger.Errorf("error audit: %v", err)
	}

	return res, err
}

// LoginSimple triggers simplified session matches and aggregates diagnostic context schemas transmitting logs straight to audit collectors.
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

// buildAudit utilizes abstract structural builders to parse identity contextual indicators, runtime trace spans, and credentials metadata.
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

// toAuditStatus inspects native error boundaries and translates execution states into unified structural telemetry indicators.
func (aaf *AuthFacadeImpl) toAuditStatus(err error) string {
	switch err == nil {
	case true:
		return auditdto.AuditStatusSuccess
	default:
		return auditdto.AuditStatusFail
	}
}
