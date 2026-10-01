package audit

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libtransport "github.com/ElfAstAhe/go-service-template/pkg/transport"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// AuthFacade extends the core facade.AuthFacade, acting as a non-invasive asynchronous broker-based auditing decorator.
// It intercepts authentication requests, structures event footprints, and broadcasts them via registered pubsub notification channels.
type AuthFacade struct {
	source    string                                      // Identifier string pinning the current microservice node instance name
	publisher pubsub.Publisher[*dto.LoginAttemptEventDTO] // Framework pubsub engine routing formatted audit events to brokers pipelines
	next      facade.AuthFacade                           // Downstream concrete authentication facade implementation handling core session logic
	logger    logger.Logger                               // Dedicated structured logging handle managing local diagnostic data
}

// Compile-time interface compliance verification
var _ facade.AuthFacade = (*AuthFacade)(nil)

// NewAuthFacadeBroker acts as a factory constructor mounting non-blocking pubsub event publishers over authentication boundaries.
func NewAuthFacadeBroker(source string, publisher pubsub.Publisher[*dto.LoginAttemptEventDTO], next facade.AuthFacade, log logger.Logger) *AuthFacade {
	return &AuthFacade{
		source:    source,
		publisher: publisher,
		next:      next,
		logger:    log,
	}
}

// Login triggers the core session scenario, aggregates runtime trace indices, and passes non-blocking audit events downstream.
func (af *AuthFacade) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// call
	res, err := af.next.Login(ctx, login)

	// audit
	data := af.buildEvent(ctx, login, res, err)

	// send
	af.publisher.Notify(ctx, data)

	return res, err
}

// LoginSimple tracks lightweight machine sessions and non-blocking broadcasts credential validation records via notification engines.
func (af *AuthFacade) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// call
	res, err := af.next.LoginSimple(ctx, login)

	// audit
	data := af.buildEvent(ctx, login, res, err)

	// send
	af.publisher.Notify(ctx, data)

	return res, err
}

// buildEvent populates the structural DTO payload tracking request sequences, client real IPs, and transaction execution errors flags.
func (af *AuthFacade) buildEvent(
	ctx context.Context,
	req *dto.LoginDTO,
	resp *dto.LoggedInDTO,
	err error,
) *dto.LoginAttemptEventDTO {
	var username = "unknown"
	if req != nil {
		username = req.Username
	}

	res := &dto.LoginAttemptEventDTO{
		NodeName:  af.source,
		EventDate: time.Now(),
		Username:  username,
		RequestID: libtransport.RequestID(ctx),
		TraceID:   libtransport.TraceID(ctx),
		IP:        libtransport.RealIP(ctx),
		Success:   true,
	}

	// Обработка ошибок
	if err != nil {
		res.Error = err.Error()
		res.Success = false
	}

	// Защита от паники на случай будущего расширения DTO:
	// Если в будущем понадобятся поля из успешного ответа (например, SessionID),
	// проверку на nil нужно делать строго здесь:
	// if resp != nil {
	//     res.SessionID = resp.SessionID
	// }

	return res
}
