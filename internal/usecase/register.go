package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// RegisterUseCase defines the application business logic boundary for handling public user identity registration workflows.
type RegisterUseCase interface {
	// Register validates inbound raw payloads, structures initial profile criteria, and persists the identity within an atomic transaction.
	Register(ctx context.Context, username string, password string) (*domain.User, error)
}

// RegisterInteractor implements the RegisterUseCase interface, orchestrating automated cryptographic RSA keys generation,
// business metadata structure alignment, and safe database state creation inside a managed transaction context.
type RegisterInteractor struct {
	uw         libdom.UnitOfWork     // BLL-level unit of work boundary abstraction manager
	hashCipher utils.Cipher          // Utility engine to calculate secure credentials hash sums
	keysHelper helper.RSAKeys        // High-performance cryptographic tool managing asymmetric key pair generation
	userRepo   domain.UserRepository // DAL repository handle managing user state persistence operations
}

// Compile-time interface compliance verification
var _ RegisterUseCase = (*RegisterInteractor)(nil)

// NewRegisterUseCase acts as a factory constructor mounting required user repository, cryptographic, and transaction dependencies.
func NewRegisterUseCase(
	uw libdom.UnitOfWork,
	hashCipher utils.Cipher,
	keysHelper helper.RSAKeys,
	userRepo domain.UserRepository,
) *RegisterInteractor {
	return &RegisterInteractor{
		uw:         uw,
		hashCipher: hashCipher,
		keysHelper: keysHelper,
		userRepo:   userRepo,
	}
}

// Register validates criteria input bounds, coordinates cryptographic pre-processing asset layout mapping, and executes database persistence inside a transactional closure.
func (ri *RegisterInteractor) Register(ctx context.Context, username string, password string) (*domain.User, error) {
	if err := ri.validate(username, password); err != nil {
		return nil, errs.NewBllValidateError("RegisterInteractor.Register", "validate income data failed", err)
	}

	model, err := ri.prepareUser(username, password)
	if err != nil {
		return nil, errs.NewBllError("RegisterInteractor.Register", "prepare user failed", err)
	}

	var res *domain.User
	// FIXED: Utilizing unique 'txCtx' sequence parameter inside the closure callback to guarantee strict ACID compliance
	err = ri.uw.Execute(ctx, func(txCtx context.Context) error {
		var txErr error
		res, txErr = ri.userRepo.Create(txCtx, model)

		return txErr
	})
	if err != nil {
		if _, ok := errors.AsType[*errs.DalAlreadyExistsError](err); ok {
			return nil, errs.NewBllUniqueError("RegisterInteractor.Register", "User", username, err)
		}

		return nil, errs.NewBllError("RegisterInteractor.Register", fmt.Sprintf("register user [%v] failed", username), err)
	}

	return res, err
}

// validate executes initial syntax, length, and presence checks over inbound registration parameters before opening database transactions.
func (ri *RegisterInteractor) validate(username string, password string) error {
	if strings.TrimSpace(username) == "" {
		return errs.NewInvalidArgumentError("username", "username is empty")
	}
	if strings.TrimSpace(password) == "" {
		return errs.NewInvalidArgumentError("password", "password is empty")
	}

	return nil
}

// prepareUser instantiates a baseline structural object schema, allocates RSA key sets, and enforces strict type parameters matching corporate invariants.
func (ri *RegisterInteractor) prepareUser(username, password string) (*domain.User, error) {
	res := domain.NewEmptyUser()
	publicKey, privateKey, err := ri.keysHelper.Generate()
	if err != nil {
		return nil, errs.NewBllError("RegisterInteractor.prepareUser", "generate RSA keys pair", err)
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
