package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	usecase "github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type ChangeKeysUseCase interface {
	ChangeKeys(ctx context.Context, userID string) (privateKey string, publicKey string, err error)
}

type ChangeKeysInteractor struct {
	keysHelper helper.RSAKeys
	tm         usecase.TransactionManager
	userRepo   domain.UserRepository
}

func NewChangeKeysUseCase(keysHelper helper.RSAKeys, tm usecase.TransactionManager, userRepo domain.UserRepository) *ChangeKeysInteractor {
	return &ChangeKeysInteractor{
		keysHelper: keysHelper,
		tm:         tm,
		userRepo:   userRepo,
	}
}

func (ck *ChangeKeysInteractor) ChangeKeys(ctx context.Context, userID string) (string, string, error) {
	if err := ck.validate(userID); err != nil {
		return "", "", domerrs.NewBllValidateError("ChangeKeysInteractor.ChangeKeys", "validate income data failed", err)
	}

	var privateKey, publicKey string

	err := ck.tm.WithinTransaction(ctx, nil, func(txCtx context.Context) error {
		// пользователь
		user, err := ck.userRepo.Find(txCtx, userID)
		if err != nil {
			return err
		}
		// генерируем пару
		privateKey, publicKey, err = ck.keysHelper.Generate()
		if err != nil {
			return domerrs.NewBllError("ChangeKeysInteractor.ChangeKeys", "generate new RSA keys", err)
		}

		// storing
		user.PrivateKey = privateKey
		user.PublicKey = publicKey

		_, err = ck.userRepo.Change(ctx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return "", "", domerrs.NewBllNotFoundError("ChangeKeysInteractor.ChangeKeys", "User", userID, err)
		}

		return "", "", domerrs.NewBllError("ChangeKeysInteractor.ChangeKeys", fmt.Sprintf("change user id [%v] keys failed", userID), err)
	}

	return privateKey, publicKey, nil
}

func (ck *ChangeKeysInteractor) validate(userID string) error {
	if strings.TrimSpace(userID) == "" {
		return errs.NewInvalidArgumentError("userID", "user id required")
	}

	return nil
}
