package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// ChangePasswordUseCase defines the business contract for validating and updating account security credentials.
type ChangePasswordUseCase interface {
	// ChangePassword modifies the password hash for the specified user identity after verifying current ownership.
	ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error
}

// ChangePasswordInteractor implements the ChangePasswordUseCase interface, orchestrating security verification workflows
// and user state modifications within an atomic transaction boundary.
type ChangePasswordInteractor struct {
	hashCipher utils.Cipher          // Cryptographic utility for secure text hashing and encryption
	uw         libdom.UnitOfWork     // BLL-level unit of work boundary abstraction manager
	userRepo   domain.UserRepository // DAL repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ ChangePasswordUseCase = (*ChangePasswordInteractor)(nil)

// NewChangePasswordUseCase acts as a factory constructor mounting interactor dependencies.
func NewChangePasswordUseCase(
	hashCipher utils.Cipher,
	uw libdom.UnitOfWork,
	userRepo domain.UserRepository,
) *ChangePasswordInteractor {
	return &ChangePasswordInteractor{
		hashCipher: hashCipher,
		uw:         uw,
		userRepo:   userRepo,
	}
}

// ChangePassword validates basic constraints, verifies current credentials, hashes the new password payload, and updates user state inside an atomic transaction.
func (cp *ChangePasswordInteractor) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	if err := cp.validate(userID, oldPassword, newPassword); err != nil {
		return errs.NewBllValidateError("ChangePasswordInteractor.ChangePassword", "validate income data failed", err)
	}

	err := cp.uw.Execute(ctx, func(txCtx context.Context) error {
		// пользователь
		user, err := cp.userRepo.Find(txCtx, userID)
		if err != nil {
			return err
		}
		// хэш сумма новый пароль
		newPasswordHash, err := cp.hashCipher.EncryptString(newPassword)
		if err != nil {
			return errs.NewBllError("ChangePasswordInteractor.ChangePassword", "new password hash build failed", err)
		}
		// хэш сумма старый пароль
		oldPasswordHash, err := cp.hashCipher.EncryptString(oldPassword)
		if err != nil {
			return errs.NewBllError("ChangePasswordInteractor.ChangePassword", "old password hash build failed", err)
		}
		// проверки
		err = cp.validatePassword(oldPasswordHash, newPasswordHash, user)
		if err != nil {
			return err
		}

		// storing
		user.PasswordHash = newPasswordHash
		user.UpdatedAt = time.Now()

		_, err = cp.userRepo.Change(txCtx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return errs.NewBllNotFoundError("ChangePasswordInteractor.ChangePassword", "User", userID, err)
		}

		return errs.NewBllError("ChangePasswordInteractor.ChangePassword", fmt.Sprintf("user id [%v] change password failed", userID), err)
	}

	return nil
}

// validate executes initial syntax and sanity checks on boundary parameters before opening a transaction.
func (cp *ChangePasswordInteractor) validate(userID, oldPassword, newPassword string) error {
	if strings.TrimSpace(userID) == "" {
		return errs.NewInvalidArgumentError("userID", "user id required")
	}
	// * empty
	if strings.TrimSpace(newPassword) == "" {
		return errs.NewInvalidArgumentError("newPassword", "new password required")
	}
	// * empty
	if strings.TrimSpace(oldPassword) == "" {
		return errs.NewInvalidArgumentError("oldPassword", "old password required")
	}

	return nil
}

// validatePassword evaluates business domain invariants regarding password history, compatibility, and credentials verification.
func (cp *ChangePasswordInteractor) validatePassword(oldPasswordHash, newPasswordHash string, user *domain.User) error {
	// * same password
	if newPasswordHash == user.PasswordHash {
		return errs.NewBllValidateError("ChangePasswordInteractor.validatePassword", "new password same as current old password", nil)
	}
	// * old and current password match
	if oldPasswordHash != user.PasswordHash {
		return errs.NewBllValidateError("ChangePasswordInteractor.validatePassword", "old password does not match current password", nil)
	}

	return nil
}
