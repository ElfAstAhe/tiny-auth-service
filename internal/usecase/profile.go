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

	res, err := p.userRepo.FindByName(ctx, username)
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return nil, domerrs.NewBllNotFoundError("ProfileUseCase.Get", "User", username, err)
		}

		return nil, domerrs.NewBllError("ProfileUseCase.Get", fmt.Sprintf("get user name [%v] failed", username), err)
	}

	return res, nil
}
