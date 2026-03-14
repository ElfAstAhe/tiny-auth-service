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
	"github.com/golang-jwt/jwt/v5"
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

// Login — основная точка входа в UseCase аутентификации.
// Выполняет валидацию входных данных, поиск пользователя, расшифровку пароля через RSA,
// проверку хэша и генерацию пары токенов (Access и Refresh).
//
// Параметры:
//   - ctx: контекст выполнения запроса.
//   - username: имя пользователя (логин).
//   - encryptedPassword: пароль, зашифрованный на публичном ключе пользователя (Base64 RSA).
//
// ToDo: переделать передачу пароля через []byte
func (luc *LoginUseCase) Login(ctx context.Context, username, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error) {
	// fails-fast
	if err := luc.validate(username, encryptedPassword); err != nil {
		return nil, nil, domerrs.NewBllValidateError("LoginUseCase.Login", fmt.Sprintf("username [%s], password [censored], invalid income data", username), err)
	}
	// пользователь
	user, err := luc.userRepo.FindByName(ctx, username)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginUseCase.Login", "load user", err)
	}
	// password hash
	passwordHash, err := luc.buildPasswordHash(user, encryptedPassword)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginUseCase.Login", "hash password", err)
	}
	// проверка пароля
	err = luc.validateUserAndPassword(user, passwordHash)
	if err != nil {
		return nil, nil, domerrs.NewBllValidateError("LoginUseCase.Login", fmt.Sprintf("user [%s], invalid credentials", username), err)
	}

	return luc.buildAnswer(user)
}

// validate выполняет первичную проверку входных параметров на пустоту (Fail-Fast).
//
// ToDo: переделать передачу пароля через []byte
func (luc *LoginUseCase) validate(username, encryptedPassword string) error {
	if strings.TrimSpace(username) == "" {
		return domerrs.NewBllValidateError("LoginUseCase.validate", "username is empty", nil)
	}
	if strings.TrimSpace(encryptedPassword) == "" {
		return domerrs.NewBllValidateError("LoginUseCase.validate", "encryptedPassword is empty", nil)
	}

	return nil
}

// buildPasswordHash расшифровывает полученный пароль с помощью приватного ключа пользователя
// и вычисляет его хэш-сумму для последующего сравнения.
//
// Параметры:
//   - user: объект доменной модели пользователя с данными о ключах и сохраненном хэше [domain.User].
//   - encryptedPassword: зашифрованная строка пароля.
//
// ToDo: переделать передачу пароля через []byte
func (luc *LoginUseCase) buildPasswordHash(user *domain.User, encryptedPassword string) (string, error) {
	// private RSA
	userPrivateKey, err := luc.keysHelper.ParsePrivateKey(user.PrivateKey)
	if err != nil {
		return "", domerrs.NewBllError("LoginUseCase.buildPasswordHash", "parse private key", err)
	}
	// decrypt password
	password, err := luc.keysHelper.DecryptString(encryptedPassword, userPrivateKey)
	if err != nil {
		return "", domerrs.NewBllError("LoginUseCase.buildPasswordHash", "decrypt password", err)
	}
	// password hash
	passwordHash, err := luc.hashHelper.EncryptString(password)
	if err != nil {
		return "", domerrs.NewBllError("LoginUseCase.buildPasswordHash", "hash password", err)
	}

	return passwordHash, nil
}

// validateUserAndPassword проверяет состояние аккаунта [domain.User] (активен/удален)
// и соответствие вычисленного хэша пароля эталонному значению из базы данных.
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

// buildAnswer оркестрирует создание финального ответа из [domain.User], инициируя генерацию
// JWT и Refresh-токена.
func (luc *LoginUseCase) buildAnswer(user *domain.User) (*jwt.Token, *jwt.Token, error) {
	subject := ToSubject(user, nil)
	token, err := luc.buildToken(subject)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginUseCase.buildAnswer", "get token from subject", err)
	}
	refreshToken, err := luc.buildRefreshToken(user)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginUseCase.buildAnswer", "get refresh token from user", err)
	}

	return token, refreshToken, nil
}

// buildToken формирует стандартный JWT Access-токен с данными пользователя и списком его ролей.
func (luc *LoginUseCase) buildToken(subject *auth.Subject) (*jwt.Token, error) {
	return luc.authHelper.TokenFromSubject(subject)
}

// buildRefreshToken генерирует уникальный токен обновления (Session-based)
// и сохраняет его состояние в хранилище сессий.
func (luc *LoginUseCase) buildRefreshToken(user *domain.User) (*jwt.Token, error) {
	// ToDo: реализовать в будущем :-)
	// ..

	return (*jwt.Token)(nil), nil
}
