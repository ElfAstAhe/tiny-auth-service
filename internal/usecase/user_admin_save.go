package usecase

import (
	"context"
	"errors"
	"fmt"

	usecase "github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type UserAdminSaveUseCase interface {
	Save(ctx context.Context, model *domain.User) (*domain.User, error)
}

type UserAdminSaveInteractor struct {
	tm         usecase.TransactionManager
	keysHelper helper.RSAKeys
	hashCipher utils.Cipher
	userRepo   domain.UserAdminRepository
}

func NewUserAdminSaveUseCase(tm usecase.TransactionManager, hashCipher utils.Cipher, keysHelper helper.RSAKeys, userRepo domain.UserAdminRepository) *UserAdminSaveInteractor {
	return &UserAdminSaveInteractor{
		tm:         tm,
		keysHelper: keysHelper,
		hashCipher: hashCipher,
		userRepo:   userRepo,
	}
}

func (uas *UserAdminSaveInteractor) Save(ctx context.Context, model *domain.User) (*domain.User, error) {
	var res *domain.User
	var err error
	// генерируем ключи
	if model.PublicKey == "" || model.PrivateKey == "" {
		model.PrivateKey, model.PublicKey, err = uas.keysHelper.Generate()
		if err != nil {
			return nil, domerrs.NewBllError("UserAdminSaveInteractor.Save", "generate new RSA keys failed", err)
		}
	}
	//// пароль для нового пользователя
	//if !model.IsExists() {
	//    model.PasswordHash, err = uas.hashCipher.EncryptString(model.PasswordHash)
	//    if err != nil {
	//        return nil, domerrs.NewBllError("UserAdminSaveInteractor.Save", "build password hash failed", err)
	//    }
	//}

	// сохраняем
	err = uas.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		var txErr error
		if !model.IsExists() {
			res, txErr = uas.userRepo.Create(ctx, model)
		} else {
			res, txErr = uas.userRepo.Change(ctx, model)
		}

		return txErr
	})
	if err != nil {
		if errors.As(err, new(*errs.DalNotFoundError)) {
			return nil, domerrs.NewBllNotFoundError("UserAdminSaveInteractor.Save", "User", model.ID, err)
		}
		if errors.As(err, new(*errs.DalAlreadyExistsError)) {
			return nil, domerrs.NewBllUniqueError("UserAdminSaveInteractor.Save", "User", model.ID, err)
		}

		return nil, domerrs.NewBllError("UserAdminSaveInteractor.Save", fmt.Sprintf("save User model id [%v] failed", model.ID), err)
	}

	return res, err
}
