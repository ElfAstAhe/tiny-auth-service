package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// ProfileUseCase defines the application business logic boundary for handling public profile lookups.
type ProfileUseCase interface {
	// Get retrieves a detailed user entity identity framework matching the provided unique username.
	Get(ctx context.Context, username string) (*domain.User, error)
}

// ProfileInteractor implements the ProfileUseCase interface, orchestrating criteria evaluation
// and data storage lookup queries to deliver identity payloads.
type ProfileInteractor struct {
	userRepo domain.UserRepository // DAL repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ ProfileUseCase = (*ProfileInteractor)(nil)

// NewProfileUseCase acts as a factory constructor mounting required user repository dependencies.
func NewProfileUseCase(userRepo domain.UserRepository) *ProfileInteractor {
	return &ProfileInteractor{
		userRepo: userRepo,
	}
}

// Get validates inbound payloads and delegates execution to the persistence layer to fetch user data aggregate fields.
func (p *ProfileInteractor) Get(ctx context.Context, username string) (*domain.User, error) {
	if err := p.validate(username); err != nil {
		return nil, errs.NewBllValidateError("ProfileInteractor.Get", "validate income data failed", err)
	}

	res, err := p.userRepo.FindByName(ctx, username)
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("ProfileInteractor.Get", "User", username, err)
		}

		return nil, errs.NewBllError("ProfileInteractor.Get", fmt.Sprintf("get user name [%v] failed", username), err)
	}

	return res, nil
}

// validate executes initial semantic text analysis over inbound query parameters before hitting storage.
func (p *ProfileInteractor) validate(username string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "must not be empty")
	}

	return nil
}
