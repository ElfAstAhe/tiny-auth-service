package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/golang-jwt/jwt/v5"
)

type LoginSimpleUseCase interface {
	Login(ctx context.Context, username string, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error)
}

type LoginSimpleInteractor struct {
	hashCipher utils.Cipher
	authHelper auth.Helper
	userRepo   domain.UserRepository
	// нотификация о логине пользователя (например аудит)
	// ...
}

var _ LoginSimpleUseCase = (*LoginSimpleInteractor)(nil)

// NewLoginSimpleUseCase создаёт новый экземпляр use case для аутентификации пользователя
//
// Параметры:
//   - hashCipher: помощник для работы с hash
//   - keysHelper: помощник для работы с RSA ключами
//   - authHelper: логика генерации токенов
//   - userRepo: репозиторий для доступа к данным пользователя
func NewLoginSimpleUseCase(hashCipher utils.Cipher, authHelper auth.Helper, userRepo domain.UserRepository) *LoginSimpleInteractor {
	return &LoginSimpleInteractor{
		hashCipher: hashCipher,
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
func (lsi *LoginSimpleInteractor) Login(ctx context.Context, username, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error) {
	// fails-fast
	if err := lsi.validate(username, encryptedPassword); err != nil {
		return nil, nil, domerrs.NewBllValidateError("LoginSimpleInteractor.Login", "validate income data failed", err)
	}
	// пользователь
	user, err := lsi.userRepo.FindByName(ctx, username)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginSimpleInteractor.Login", "load user", err)
	}
	// password hash
	passwordHash, err := lsi.buildPasswordHash(user, encryptedPassword)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginSimpleInteractor.Login", "hash password", err)
	}
	// проверка пароля
	err = lsi.validateUserAndPassword(user, passwordHash)
	if err != nil {
		return nil, nil, domerrs.NewBllValidateError("LoginSimpleInteractor.Login", fmt.Sprintf("user [%s], invalid credentials", username), err)
	}

	return lsi.buildAnswer(user)
}

// validate выполняет первичную проверку входных параметров на пустоту (Fail-Fast).
//
// ToDo: переделать передачу пароля через []byte
func (lsi *LoginSimpleInteractor) validate(username, encryptedPassword string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "username is required")
	}
	if strings.TrimSpace(encryptedPassword) == "" {
		return errs.NewInvalidArgumentError("encryptedPassword", "encrypted password is required")
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
func (lsi *LoginSimpleInteractor) buildPasswordHash(user *domain.User, password string) (string, error) {
	// password hash
	passwordHash, err := lsi.hashCipher.EncryptString(password)
	if err != nil {
		return "", domerrs.NewBllError("LoginSimpleInteractor.buildPasswordHash", "hash password", err)
	}

	return passwordHash, nil
}

// validateUserAndPassword проверяет состояние аккаунта [domain.User] (активен/удален)
// и соответствие вычисленного хэша пароля эталонному значению из базы данных.
func (lsi *LoginSimpleInteractor) validateUserAndPassword(user *domain.User, passwordHash string) error {
	// active
	if !user.Active {
		return domerrs.NewBllUnauthorizedError("LoginSimpleInteractor.validateUserAndPassword", "user is not active", nil)
	}
	// deleted
	if user.Deleted {
		return domerrs.NewBllUnauthorizedError("LoginSimpleInteractor.validateUserAndPassword", "user is deleted", nil)
	}
	// passwords
	if user.PasswordHash != passwordHash {
		return domerrs.NewBllUnauthorizedError("LoginSimpleInteractor.validateUserAndPassword", "user password hash is invalid", nil)
	}

	return nil
}

// buildAnswer оркестрирует создание финального ответа из [domain.User], инициируя генерацию
// JWT и Refresh-токена.
func (lsi *LoginSimpleInteractor) buildAnswer(user *domain.User) (*jwt.Token, *jwt.Token, error) {
	subject := ToSubject(user, nil)
	token, err := lsi.buildToken(subject)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginSimpleInteractor.buildAnswer", "build token from subject", err)
	}
	refreshToken, err := lsi.buildRefreshToken(user)
	if err != nil {
		return nil, nil, domerrs.NewBllError("LoginSimpleInteractor.buildAnswer", "build refresh token from user", err)
	}

	return token, refreshToken, nil
}

// buildToken формирует стандартный JWT Access-токен с данными пользователя и списком его ролей.
func (lsi *LoginSimpleInteractor) buildToken(subject *auth.Subject) (*jwt.Token, error) {
	return lsi.authHelper.TokenFromSubject(subject)
}

// buildRefreshToken генерирует уникальный токен обновления (Session-based)
// и сохраняет его состояние в хранилище сессий.
func (lsi *LoginSimpleInteractor) buildRefreshToken(user *domain.User) (*jwt.Token, error) {
	// ToDo: реализовать в будущем :-)
	// ..

	return (*jwt.Token)(nil), nil
}
