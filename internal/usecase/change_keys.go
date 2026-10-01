package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// ChangeKeysUseCase defines the business contract for generating and rotating asymmetric identity credentials.
type ChangeKeysUseCase interface {
	// ChangeKeys executes the core cryptographic key pair rotation for the specified user identity.
	ChangeKeys(ctx context.Context, userID string) (privateKey string, publicKey string, err error)
}

// ChangeKeysInteractor implements the ChangeKeysUseCase interface, orchestrating core business scenarios
// for generating new RSA keys and updating data assets inside a managed transaction context.
type ChangeKeysInteractor struct {
	keysHelper helper.RSAKeys        // High-performance cryptographic key generation utility
	uw         libdom.UnitOfWork     // BLL-level unit of work boundary abstraction manager
	userRepo   domain.UserRepository // DAL repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ ChangeKeysUseCase = (*ChangeKeysInteractor)(nil)

// NewChangeKeysUseCase acts as a factory constructor mounting interactor dependencies.
func NewChangeKeysUseCase(
	keysHelper helper.RSAKeys,
	uw libdom.UnitOfWork,
	userRepo domain.UserRepository,
) *ChangeKeysInteractor {
	return &ChangeKeysInteractor{
		keysHelper: keysHelper,
		uw:         uw,
		userRepo:   userRepo,
	}
}

// ChangeKeys validates input criteria and executes full RSA generation and storage sequence wrapped in an atomic boundary.
func (ck *ChangeKeysInteractor) ChangeKeys(ctx context.Context, userID string) (string, string, error) {
	if err := ck.validate(userID); err != nil {
		return "", "", errs.NewBllValidateError("ChangeKeysInteractor.ChangeKeys", "validate income data failed", err)
	}

	var privateKey, publicKey string

	err := ck.uw.Execute(ctx, func(txCtx context.Context) error {
		// пользователь
		user, err := ck.userRepo.Find(txCtx, userID)
		if err != nil {
			return err
		}
		// генерируем пару
		privateKey, publicKey, err = ck.keysHelper.Generate()
		if err != nil {
			return errs.NewBllError("ChangeKeysInteractor.ChangeKeys", "generate new RSA keys", err)
		}

		// storing
		user.PrivateKey = privateKey
		user.PublicKey = publicKey

		// FIXED: Utilizing the active transactional 'txCtx' to guarantee strict ACID atomic boundary isolation
		_, err = ck.userRepo.Change(txCtx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return "", "", errs.NewBllNotFoundError("ChangeKeysInteractor.ChangeKeys", "User", userID, err)
		}

		return "", "", errs.NewBllError("ChangeKeysInteractor.ChangeKeys", fmt.Sprintf("change user id [%v] keys failed", userID), err)
	}

	return privateKey, publicKey, nil
}

// validate executes semantic syntax analysis over input criteria payloads.
func (ck *ChangeKeysInteractor) validate(userID string) error {
	if strings.TrimSpace(userID) == "" {
		return errs.NewInvalidArgumentError("userID", "user id required")
	}

	return nil
}
