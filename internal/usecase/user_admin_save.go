package usecase

import (
	"context"
	"errors"
	"fmt"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAdminSaveUseCase defines the application business logic boundary for handling administrative user persistence and modification workflows.
type UserAdminSaveUseCase interface {
	// Save creates a new user identity or modifies an existing one, automatically ensuring cryptographic key pair provisioning.
	Save(ctx context.Context, model *domain.User) (*domain.User, error)
}

// UserAdminSaveInteractor implements the UserAdminSaveUseCase interface, orchestrating automatic cryptographic RSA keys provisioning,
// dynamic operational routing between creation and updates, and ACID state mapping within a unit of work context.
type UserAdminSaveInteractor struct {
	uw         libdom.UnitOfWork          // BLL-level unit of work boundary abstraction manager
	keysHelper helper.RSAKeys             // High-performance cryptographic utility for automated asymmetric keys generation
	hashCipher utils.Cipher               // Utility engine to calculate secure user credentials hash sums
	userRepo   domain.UserAdminRepository // DAL administrative repository handle for user state persistence
}

// Compile-time interface compliance verification
var _ UserAdminSaveUseCase = (*UserAdminSaveInteractor)(nil)

// NewUserAdminSaveUseCase acts as a factory constructor mounting required user administration repository, cryptographic, and transaction dependencies.
func NewUserAdminSaveUseCase(
	uw libdom.UnitOfWork,
	hashCipher utils.Cipher,
	keysHelper helper.RSAKeys,
	userRepo domain.UserAdminRepository,
) *UserAdminSaveInteractor {
	return &UserAdminSaveInteractor{
		uw:         uw,
		keysHelper: keysHelper,
		hashCipher: hashCipher,
		userRepo:   userRepo,
	}
}

// Save evaluates the presence of identity key sets, triggers automated cryptographic asset provisioning if missing, and commits structural states inside a transactional closure.
func (uas *UserAdminSaveInteractor) Save(ctx context.Context, model *domain.User) (*domain.User, error) {
	var res *domain.User
	var err error
	// генерируем ключи
	if model.PublicKey == "" || model.PrivateKey == "" {
		model.PrivateKey, model.PublicKey, err = uas.keysHelper.Generate()
		if err != nil {
			return nil, errs.NewBllError("UserAdminSaveInteractor.Save", "generate new RSA keys failed", err)
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
	// FIXED: Utilizing unique 'txCtx' sequence parameter inside the closure callback to guarantee strict ACID compliance
	err = uas.uw.Execute(ctx, func(txCtx context.Context) error {
		var txErr error
		if !model.IsExists() {
			res, txErr = uas.userRepo.Create(txCtx, model)
		} else {
			res, txErr = uas.userRepo.Change(txCtx, model)
		}

		return txErr
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalNotFoundError](err); ok {
			return nil, errs.NewBllNotFoundError("UserAdminSaveInteractor.Save", "User", model.ID, err)
		}
		if _, ok := errors.AsType[*errs.DalAlreadyExistsError](err); ok {
			return nil, errs.NewBllUniqueError("UserAdminSaveInteractor.Save", "User", model.ID, err)
		}

		return nil, errs.NewBllError("UserAdminSaveInteractor.Save", fmt.Sprintf("save User model id [%v] failed", model.ID), err)
	}

	return res, err
}
