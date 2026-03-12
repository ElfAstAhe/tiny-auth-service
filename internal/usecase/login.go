package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type LoginUseCase struct {
	hashHelper utils.Cipher
	keysHelper *helper.RSAKeysHelper
	authHelper *auth.Helper
	userRepo   domain.UserRepository
	// нотификация о логине пользователя (например аудит)
	// ...
}

// NewLoginUseCase создаёт новый экземпляр use case для аутентификации пользователя
//
// Параметры:
//   - hashHelper: помощник для работы с hash
//   - keysHelper: помощник для работы с RSA ключами
//   - authHelper: логика генерации токенов
//   - userRepo: репозиторий для доступа к данным пользователя
func NewLoginUseCase(hashHelper utils.Cipher, keysHelper *helper.RSAKeysHelper, authHelper *auth.Helper, userRepo domain.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		hashHelper: hashHelper,
		keysHelper: keysHelper,
		authHelper: authHelper,
		userRepo:   userRepo,
	}
}

func (luc *LoginUseCase) Login(ctx context.Context, username, encryptedPassword string) (token string, refreshToken string, err error) {
	// fails-fast
	if err := luc.validate(username, encryptedPassword); err != nil {
		return "", "", domerrs.NewBllValidateError("LoginUseCase.Login", fmt.Sprintf("username [%s], password [censored], invalid income data", username), err)
	}
	// пользователь
	user, err := luc.userRepo.FindByName(ctx, username)
	if err != nil {
		return "", "", domerrs.NewBllError("LoginUseCase.Login", "load user", err)
	}
	// password hash
	passwordHash, err := luc.preparePasswordHash(user, encryptedPassword)
	if err != nil {
		return "", "", domerrs.NewBllError("LoginUseCase.Login", "hash password", err)
	}
	// проверка пароля
	err = luc.validateUserAndPassword(user, passwordHash)
	if err != nil {
		return "", "", domerrs.NewBllValidateError("LoginUseCase.Login", fmt.Sprintf("user [%s], invalid credentials", username), err)
	}

	return luc.buildAnswer(user)
}

func (luc *LoginUseCase) validate(username, encryptedPassword string) error {
	if strings.TrimSpace(username) == "" {
		return domerrs.NewBllValidateError("LoginUseCase.validate", "username is empty", nil)
	}
	if strings.TrimSpace(encryptedPassword) == "" {
		return domerrs.NewBllValidateError("LoginUseCase.validate", "encryptedPassword is empty", nil)
	}

	return nil
}

func (luc *LoginUseCase) preparePasswordHash(user *domain.User, encryptedPassword string) (string, error) {
	// private RSA
	userPrivateKey, err := luc.keysHelper.ParsePrivateKey(user.PrivateKey)
	if err != nil {
		return "", domerrs.NewBllError("LoginUseCase.preparePasswordHash", "parse private key", err)
	}
	// decrypt password
	password, err := luc.keysHelper.DecryptString(encryptedPassword, userPrivateKey)
	if err != nil {
		return "", domerrs.NewBllError("LoginUseCase.preparePasswordHash", "decrypt password", err)
	}
	// password hash
	passwordHash, err := luc.hashHelper.EncryptString(password)
	if err != nil {
		return "", domerrs.NewBllError("LoginUseCase.preparePasswordHash", "hash password", err)
	}

	return passwordHash, nil
}

func (luc *LoginUseCase) validateUserAndPassword(user *domain.User, passwordHash string) error {
	// active
	if !user.Active {
		return domerrs.NewBllUnauthorizedError("LoginUseCase.validateUserAndPassword", "user is not active", nil)
	}
	// deleted
	if user.Deleted {
		return domerrs.NewBllUnauthorizedError("LoginUseCase.validateUserAndPassword", "user is deleted", nil)
	}
	// passwords
	if user.PasswordHash != passwordHash {
		return domerrs.NewBllUnauthorizedError("LoginUseCase.validateUserAndPassword", "user password hash is invalid", nil)
	}

	return nil
}

func (luc *LoginUseCase) buildAnswer(user *domain.User) (string, string, error) {

}

func (luc *LoginUseCase) buildToken(ctx context.Context, user *domain.User) (string, error) {

}

func (luc *LoginUseCase) buildRefreshToken(ctx context.Context, user *domain.User) (string, error) {

}
