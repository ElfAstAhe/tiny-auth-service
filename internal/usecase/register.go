package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	usecase "github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type RegisterUseCase interface {
	Register(ctx context.Context, username string, password string) (*domain.User, error)
}

type RegisterInteractor struct {
	tm         usecase.TransactionManager
	hashCipher utils.Cipher
	keysHelper helper.RSAKeys
	userRepo   domain.UserRepository
}

var _ RegisterUseCase = (*RegisterInteractor)(nil)

func NewRegisterUseCase(tm usecase.TransactionManager, hashCipher utils.Cipher, keysHelper helper.RSAKeys, userRepo domain.UserRepository) *RegisterInteractor {
	return &RegisterInteractor{
		tm:         tm,
		hashCipher: hashCipher,
		keysHelper: keysHelper,
		userRepo:   userRepo,
	}
}

func (ri *RegisterInteractor) Register(ctx context.Context, username string, password string) (*domain.User, error) {
	if err := ri.validate(username, password); err != nil {
		return nil, domerrs.NewBllValidateError("RegisterInteractor.Register", "validate income data failed", err)
	}

	model, err := ri.prepareUser(username, password)
	if err != nil {
		return nil, domerrs.NewBllError("RegisterInteractor.Register", "prepare user failed", err)
	}

	var res *domain.User
	err = ri.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		var txErr error
		res, txErr = ri.userRepo.Create(ctx, model)

		return txErr
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalAlreadyExistsError](err); ok {
			return nil, domerrs.NewBllUniqueError("RegisterInteractor.Register", "User", username, err)
		}

		return nil, domerrs.NewBllError("RegisterInteractor.Register", fmt.Sprintf("register user [%v] failed", username), err)
	}

	return res, err
}

func (ri *RegisterInteractor) validate(username string, password string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "username is empty")
	}
	if strings.TrimSpace(password) == "" {
		return errs.NewInvalidArgumentError("password", "password is empty")
	}

	return nil
}

func (ri *RegisterInteractor) prepareUser(username, password string) (*domain.User, error) {
	res := domain.NewEmptyUser()
	publicKey, privateKey, err := ri.keysHelper.Generate()
	if err != nil {
		return nil, domerrs.NewBllError("RegisterInteractor.prepareUser", "generate RSA keys pair", err)
	}
	//passwordHash, err := ri.hashCipher.EncryptString(password)
	//if err != nil {
	//	return nil, domerrs.NewBllError("RegisterInteractor.prepareUser", "hash password failed", err)
	//}

	res.ID = ""
	res.Name = username
	// register can create only user type users (it is a business logic)
	res.Type = domain.UserTypeUser
	res.PasswordHash = password
	res.PublicKey = publicKey
	res.PrivateKey = privateKey
	res.Active = false
	res.Deleted = false
	res.CreatedAt = time.Now()
	res.UpdatedAt = time.Now()

	return res, nil
}
