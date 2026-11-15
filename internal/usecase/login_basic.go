package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

// LoginBasicUseCase defines the simplified application boundary contract for authentication workflows bypassing asymmetric encryption steps.
type LoginBasicUseCase interface {
	// Login validates credentials directly and yields cryptographically signed JWT access token pairs.
	Login(ctx context.Context, username string, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error)
}

// LoginBasicInteractor implements the LoginBasicUseCase interface, orchestrating simplified identity checking
// and raw password hashing sequences without intermediate asymmetric private key operations.
type LoginBasicInteractor struct {
	hashCipher utils.Cipher          // Utility engine to recalculate secure password hash sums
	authHelper auth.Helper           // Framework core helper orchestrating physical token structural layout packing
	userRepo   domain.UserRepository // DAL repository handle managing user identity search operations
	// нотификация о логине пользователя (например аудит)
	// ...
}

// Compile-time interface compliance verification
var _ LoginBasicUseCase = (*LoginBasicInteractor)(nil)

// NewLoginBasic creates a new LoginBasicUseCase instance wrapping direct hashing and storage infrastructure dependencies.
func NewLoginBasic(hashCipher utils.Cipher, authHelper auth.Helper, userRepo domain.UserRepository) *LoginBasicInteractor {
	return &LoginBasicInteractor{
		hashCipher: hashCipher,
		authHelper: authHelper,
		userRepo:   userRepo,
	}
}

// Login executes fast input evaluation, repository identity matching, direct encryption hashing, and final response assembly.
//
// ToDo: переделать передачу пароля через []byte
func (lbi *LoginBasicInteractor) Login(ctx context.Context, username, encryptedPassword string) (token *jwt.Token, refreshToken *jwt.Token, err error) {
	// fails-fast
	if err := lbi.validate(username, encryptedPassword); err != nil {
		return nil, nil, errs.NewBllValidateError("LoginSimpleInteractor.Login", "validate income data failed", err)
	}
	// пользователь
	user, err := lbi.userRepo.FindByName(ctx, username)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginSimpleInteractor.Login", "load user", err)
	}
	// password hash
	passwordHash, err := lbi.buildPasswordHash(user, encryptedPassword)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginSimpleInteractor.Login", "hash password", err)
	}
	// проверка пароля
	err = lbi.validateUserAndPassword(user, passwordHash)
	if err != nil {
		return nil, nil, errs.NewBllValidateError("LoginSimpleInteractor.Login", fmt.Sprintf("user [%s], invalid credentials", username), err)
	}

	return lbi.buildAnswer(user)
}

// validate executes preliminary syntax and presence checks over raw boundaries parameters.
//
// ToDo: переделать передачу пароля через []byte
func (lbi *LoginBasicInteractor) validate(username, encryptedPassword string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "username is required")
	}
	if strings.TrimSpace(encryptedPassword) == "" {
		return errs.NewInvalidArgumentError("encryptedPassword", "encrypted password is required")
	}

	return nil
}

// buildPasswordHash generates a secure comparative hash sum out of the inbound password string criteria.
//
// ToDo: переделать передачу пароля через []byte
func (lbi *LoginBasicInteractor) buildPasswordHash(user *domain.User, password string) (string, error) {
	// password hash
	passwordHash, err := lbi.hashCipher.EncryptString(password)
	if err != nil {
		return "", errs.NewBllError("LoginSimpleInteractor.buildPasswordHash", "hash password", err)
	}

	return passwordHash, nil
}

// validateUserAndPassword reviews aggregate active status fields before confirming matching password credentials state.
func (lbi *LoginBasicInteractor) validateUserAndPassword(user *domain.User, passwordHash string) error {
	// active
	if !user.Active {
		return errs.NewBllUnauthorizedError("LoginSimpleInteractor.validateUserAndPassword", "user is not active", nil)
	}
	// deleted
	if user.Deleted {
		return errs.NewBllUnauthorizedError("LoginSimpleInteractor.validateUserAndPassword", "user is deleted", nil)
	}
	// passwords
	if user.PasswordHash != passwordHash {
		return errs.NewBllUnauthorizedError("LoginSimpleInteractor.validateUserAndPassword", "user password hash is invalid", nil)
	}

	return nil
}

// buildAnswer coordinates underlying mappings converting identities into structured tokens payloads.
func (lbi *LoginBasicInteractor) buildAnswer(user *domain.User) (*jwt.Token, *jwt.Token, error) {
	subject := ToSubject(user, nil)
	token, err := lbi.buildToken(subject)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginSimpleInteractor.buildAnswer", "build token from subject", err)
	}
	refreshToken, err := lbi.buildRefreshToken(user)
	if err != nil {
		return nil, nil, errs.NewBllError("LoginSimpleInteractor.buildAnswer", "build refresh token from user", err)
	}

	return token, refreshToken, nil
}

// buildToken wraps core claims into standard signed cryptographically signed access tokens.
func (lbi *LoginBasicInteractor) buildToken(subject *auth.Subject) (*jwt.Token, error) {
	return lbi.authHelper.TokenFromSubject(subject)
}

// buildRefreshToken yields session tracking structures for persistent token rotation configurations.
func (lbi *LoginBasicInteractor) buildRefreshToken(user *domain.User) (*jwt.Token, error) {
	// ToDo: реализовать в будущем :-)
	// ..

	return (*jwt.Token)(nil), nil
}
