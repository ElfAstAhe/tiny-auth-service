package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

// LoginUseCase defines the high-level application boundary contract for handling identity authentication scenarios.
type LoginUseCase interface {
	// Login validates input parameters, decrypts securely payloads and yields cryptographically signed JWT access tokens.
	Login(ctx context.Context, username string, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error)
}

// LoginInteractor implements the LoginUseCase interface, orchestrating full identity verification sequences,
// cryptographic private RSA key decryption, credentials comparison, and downstream token generation mechanics.
type LoginInteractor struct {
	hashCipher utils.Cipher          // Utility engine to recalculate secure password hash sums
	keysHelper helper.RSAKeys        // Cryptographic tool managing user private key data loading and transformation
	authHelper auth.Helper           // Framework core helper orchestrating physical token structural layout packing
	userRepo   domain.UserRepository // DAL repository handle managing user identity search operations
	// нотификация о логине пользователя (например аудит)
	// ...
}

// Compile-time interface compliance verification
var _ LoginUseCase = (*LoginInteractor)(nil)

// NewLoginUseCase creates a new LoginUseCase instance mounting required cryptographic and repository dependencies.
func NewLoginUseCase(hashCipher utils.Cipher, keysHelper helper.RSAKeys, authHelper auth.Helper, userRepo domain.UserRepository) *LoginInteractor {
	return &LoginInteractor{
		hashCipher: hashCipher,
		keysHelper: keysHelper,
		authHelper: authHelper,
		userRepo:   userRepo,
	}
}

// Login executes primary entry criteria checking, entity lookup, RSA decryption, hash verification, and final token response assembly.
//
// ToDo: переделать передачу пароля через []byte
func (luc *LoginInteractor) Login(ctx context.Context, username, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error) {
	// fails-fast
	if err := luc.validate(username, encryptedPassword); err != nil {
		return nil, nil, errs.NewBllValidateError("LoginInteractor.Login", "validate income data failed", err)
	}
	// пользователь
	user, err := luc.userRepo.FindByName(ctx, username)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginInteractor.Login", "load user", err)
	}
	// password hash
	passwordHash, err := luc.buildPasswordHash(user, encryptedPassword)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginInteractor.Login", "hash password", err)
	}
	// проверка пароля
	err = luc.validateUserAndPassword(user, passwordHash)
	if err != nil {
		return nil, nil, errs.NewBllValidateError("LoginInteractor.Login", fmt.Sprintf("user [%s], invalid credentials", username), err)
	}

	return luc.buildAnswer(user)
}

// validate executes initial semantic text syntax assertions over parameters to ensure strict input fail-fast enforcement.
//
// ToDo: переделать передачу пароля через []byte
func (luc *LoginInteractor) validate(username, encryptedPassword string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "username is required")
	}
	if strings.TrimSpace(encryptedPassword) == "" {
		return errs.NewInvalidArgumentError("encryptedPassword", "encrypted password is required")
	}

	return nil
}

// buildPasswordHash parses user private key structures, decodes asymmetric RSA payloads, and yields a comparative hash string.
//
// ToDo: переделать передачу пароля через []byte
func (luc *LoginInteractor) buildPasswordHash(user *domain.User, encryptedPassword string) (string, error) {
	// private RSA
	userPrivateKey, err := luc.keysHelper.ParsePrivateKey(user.PrivateKey)
	if err != nil {
		return "", errs.NewBllError("LoginInteractor.buildPasswordHash", "parse private key", err)
	}
	// decrypt password
	password, err := luc.keysHelper.DecryptString(encryptedPassword, userPrivateKey)
	if err != nil {
		return "", errs.NewBllError("LoginInteractor.buildPasswordHash", "decrypt password", err)
	}
	// password hash
	passwordHash, err := luc.hashCipher.EncryptString(password)
	if err != nil {
		return "", errs.NewBllError("LoginInteractor.buildPasswordHash", "hash password", err)
	}

	return passwordHash, nil
}

// validateUserAndPassword checks active and deleted account status properties before matching cryptographic comparative keys.
func (luc *LoginInteractor) validateUserAndPassword(user *domain.User, passwordHash string) error {
	// active
	if !user.Active {
		return errs.NewBllUnauthorizedError("LoginInteractor.validateUserAndPassword", "user is not active", nil)
	}
	// deleted
	if user.Deleted {
		return errs.NewBllUnauthorizedError("LoginInteractor.validateUserAndPassword", "user is deleted", nil)
	}
	// passwords
	if user.PasswordHash != passwordHash {
		return errs.NewBllUnauthorizedError("LoginInteractor.validateUserAndPassword", "user password hash is invalid", nil)
	}

	return nil
}

// buildAnswer coordinates the final data transformations mapping user identities into signed token structures.
func (luc *LoginInteractor) buildAnswer(user *domain.User) (*jwt.Token, *jwt.Token, error) {
	subject := ToSubject(user, nil)
	token, err := luc.buildToken(subject)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginInteractor.buildAnswer", "build token from subject", err)
	}
	refreshToken, err := luc.buildRefreshToken(user)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginInteractor.buildAnswer", "build refresh token from user", err)
	}

	return token, refreshToken, nil
}

// buildToken maps identity claims structures directly into standard signed JWT access identifiers.
func (luc *LoginInteractor) buildToken(subject *auth.Subject) (*jwt.Token, error) {
	return luc.authHelper.TokenFromSubject(subject)
}

// buildRefreshToken initializes session structures to assemble a secure, long-lived token rotation key.
func (luc *LoginInteractor) buildRefreshToken(user *domain.User) (*jwt.Token, error) {
	// ToDo: реализовать в будущем :-)
	// ..

	return (*jwt.Token)(nil), nil
}
