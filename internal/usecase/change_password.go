package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	usecase "github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
)

type ChangePasswordUseCase interface {
	ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error
}

type ChangePasswordInteractor struct {
	hashCipher utils.Cipher
	tm         usecase.TransactionManager
	userRepo   domain.UserRepository
}

var _ ChangePasswordUseCase = (*ChangePasswordInteractor)(nil)

func NewChangePasswordUseCase(hashCipher utils.Cipher, tm usecase.TransactionManager, userRepo domain.UserRepository) *ChangePasswordInteractor {
	return &ChangePasswordInteractor{
		hashCipher: hashCipher,
		tm:         tm,
		userRepo:   userRepo,
	}
}

func (cp *ChangePasswordInteractor) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	if err := cp.validate(userID, oldPassword, newPassword); err != nil {
		return domerrs.NewBllValidateError("ChangePasswordInteractor.ChangePassword", "validate income data failed", err)
	}

	err := cp.tm.WithinTransaction(ctx, nil, func(ctx context.Context) error {
		// пользователь
		user, err := cp.userRepo.Find(ctx, userID)
		if err != nil {
			return err
		}
		// хэш сумма новый пароль
		newPasswordHash, err := cp.hashCipher.EncryptString(newPassword)
		if err != nil {
			return domerrs.NewBllError("ChangePasswordInteractor.ChangePassword", "new password hash build failed", err)
		}
		// хэш сумма старый пароль
		oldPasswordHash, err := cp.hashCipher.EncryptString(oldPassword)
		if err != nil {
			return domerrs.NewBllError("ChangePasswordInteractor.ChangePassword", "old password hash build failed", err)
		}
		// проверки
		err = cp.validatePassword(oldPasswordHash, newPasswordHash, user)
		if err != nil {
			return err
		}

		// storing
		user.PasswordHash = newPasswordHash
		user.UpdatedAt = time.Now()

		_, err = cp.userRepo.Change(ctx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return domerrs.NewBllNotFoundError("ChangePasswordInteractor.ChangePassword", "User", userID, err)
		}

		return domerrs.NewBllError("ChangePasswordInteractor.ChangePassword", fmt.Sprintf("user id [%v] change password failed", userID), err)
	}

	return nil
}

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

// validatePassword
// сюда можно добавить бизнес логику на проверку пароля
//   - strong password
//   - same password
//   - history password
//   - min length
//   - etc
//
// реализовываем простые проверки
//   - empty
//   - same password
//   - old and current password match
func (cp *ChangePasswordInteractor) validatePassword(oldPasswordHash, newPasswordHash string, user *domain.User) error {
	// * same password
	if newPasswordHash == user.PasswordHash {
		return domerrs.NewBllValidateError("ChangePasswordInteractor.validatePassword", "new password same as current old password", nil)
	}
	// * old and current password match
	if oldPasswordHash != user.PasswordHash {
		return domerrs.NewBllValidateError("ChangePasswordInteractor.validatePassword", "old password does not match current password", nil)
	}

	return nil
}
