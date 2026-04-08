package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type ProfileUseCase interface {
	Get(ctx context.Context, username string) (*domain.User, error)
}

type ProfileInteractor struct {
	userRepo domain.UserRepository
}

var _ ProfileUseCase = (*ProfileInteractor)(nil)

func NewProfileUseCase(userRepo domain.UserRepository) *ProfileInteractor {
	return &ProfileInteractor{
		userRepo: userRepo,
	}
}

func (p *ProfileInteractor) Get(ctx context.Context, username string) (*domain.User, error) {
	if err := p.validate(username); err != nil {
		return nil, domerrs.NewBllValidateError("ProfileInteractor.Get", "validate income data failed", err)
	}

	res, err := p.userRepo.FindByName(ctx, username)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, domerrs.NewBllNotFoundError("ProfileInteractor.Get", "User", username, err)
		}

		return nil, domerrs.NewBllError("ProfileInteractor.Get", fmt.Sprintf("get user name [%v] failed", username), err)
	}

	return res, nil
}

func (p *ProfileInteractor) validate(username string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "must not be empty")
	}

	return nil
}
