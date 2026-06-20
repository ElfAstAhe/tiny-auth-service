package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type AuthFacade struct {
	publisher pubsub.Publisher[*dto.LoginAttemptEventDTO]
	next      facade.AuthFacade
	logger    logger.Logger
}

var _ facade.AuthFacade = (*AuthFacade)(nil)

func NewAuthFacadeAMQP(publisher pubsub.Publisher[*dto.LoginAttemptEventDTO], next facade.AuthFacade, log logger.Logger) *AuthFacade {
	return &AuthFacade{
		publisher: publisher,
		next:      next,
		logger:    log,
	}
}

func (a AuthFacade) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthFacade) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	//TODO implement me
	panic("implement me")
}
