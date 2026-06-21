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

type AuthFacade struct {
	source    string
	publisher pubsub.Publisher[*dto.LoginAttemptEventDTO]
	next      facade.AuthFacade
	logger    logger.Logger
}

var _ facade.AuthFacade = (*AuthFacade)(nil)

func NewAuthFacadeAMQP(source string, publisher pubsub.Publisher[*dto.LoginAttemptEventDTO], next facade.AuthFacade, log logger.Logger) *AuthFacade {
	return &AuthFacade{
		source:    source,
		publisher: publisher,
		next:      next,
		logger:    log,
	}
}

func (af *AuthFacade) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// call
	res, err := af.next.Login(ctx, login)

	// audit
	data := af.buildEvent(ctx, login, res, err)

	// send
	af.publisher.Notify(ctx, data)

	return res, err
}

func (af *AuthFacade) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	// call
	res, err := af.next.LoginSimple(ctx, login)

	// audit
	data := af.buildEvent(ctx, login, res, err)

	// send
	af.publisher.Notify(ctx, data)

	return res, err
}

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
