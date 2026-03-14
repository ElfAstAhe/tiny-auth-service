package usecase

import (
	"context"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type ProfileUseCase struct {
	userRepo domain.UserRepository
}

func NewProfileUseCase(userRepo domain.UserRepository) *ProfileUseCase {
	return &ProfileUseCase{
		userRepo: userRepo,
	}
}

func (p *ProfileUseCase) Get(ctx context.Context, username string) (*domain.User, error) {
	if strings.TrimSpace(username) == "" {
		return nil, errs.NewInvalidArgumentError("username", "must not be empty")
	}

	return p.buildAnswer(ctx, username)
}

func (p *ProfileUseCase) buildAnswer(ctx context.Context, username string) (*domain.User, error) {

}
