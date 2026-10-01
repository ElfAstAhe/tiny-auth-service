package facade

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

// UserFacade defines the public user domain orchestration contract managing data mapping, profile querying, and identity mutations.
type UserFacade interface {
	Register(ctx context.Context, register *dto.RegisterDTO) (*dto.ProfileDTO, error)
	Profile(ctx context.Context) (*dto.ProfileDTO, error)
	ChangePassword(ctx context.Context, changePassword *dto.ChangePasswordDTO) error
	ChangeKeys(ctx context.Context) (*dto.ChangedKeysDTO, error)
}

// UserFacadeImpl structures public boundary routers, conducting data transformations and non-admin role assertions.
type UserFacadeImpl struct {
	authHelper       auth.Helper                   // Framework core token management tool verifying structural context identity claims
	registerUC       usecase.RegisterUseCase       // Downstream usecase executing public registration and initial asset provisioning workflows
	profileUC        usecase.ProfileUseCase        // Downstream usecase managing fine-grained profile lookup and retrieval scenarios
	changePasswordUC usecase.ChangePasswordUseCase // Downstream usecase coordinating credentials verification and password mutation boundaries
	changeKeysUC     usecase.ChangeKeysUseCase     // Downstream usecase managing asymmetric key rotation and cryptographic operations
}

// Compile-time interface compliance verification
var _ UserFacade = (*UserFacadeImpl)(nil)

// NewUserFacade acts as a factory constructor embedding modular user-facing usecase components and framework helpers.
func NewUserFacade(
	authHelper auth.Helper,
	registerUC usecase.RegisterUseCase,
	profileUC usecase.ProfileUseCase,
	changePasswordUC usecase.ChangePasswordUseCase,
	changeKeysUC usecase.ChangeKeysUseCase,
) *UserFacadeImpl {
	return &UserFacadeImpl{
		authHelper:       authHelper,
		registerUC:       registerUC,
		profileUC:        profileUC,
		changePasswordUC: changePasswordUC,
		changeKeysUC:     changeKeysUC,
	}
}

// Register forwards credentials data down to usecases and encapsulates the returned domain user instance within a transmittable ProfileDTO.
func (uf *UserFacadeImpl) Register(ctx context.Context, register *dto.RegisterDTO) (*dto.ProfileDTO, error) {
	res, err := uf.registerUC.Register(ctx, register.Username, register.Password)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToProfileDTO(res), nil
}

// Profile extracts token context claims, asserts role invariants, and yields a mapped data transfer object detailing user properties.
func (uf *UserFacadeImpl) Profile(ctx context.Context) (*dto.ProfileDTO, error) {
	// subject
	subj, err := uf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserFacadeImpl.Profile", "retrieve subject", err)
	}
	// rbac
	if !subj.HasRole(domain.RoleUser) {
		return nil, errs.NewBllForbiddenError("UserFacadeImpl.Profile", "user is not a user", nil)
	}
	// logic
	res, err := uf.profileUC.Get(ctx, subj.Name)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToProfileDTO(res), nil
}

// ChangePassword evaluates contextual subject tokens and enforces strict updates mapping new credential hashes inside domain stores.
func (uf *UserFacadeImpl) ChangePassword(ctx context.Context, changePassword *dto.ChangePasswordDTO) error {
	// subject
	subj, err := uf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return errs.NewBllForbiddenError("UserFacadeImpl.ChangePassword", "retrieve subject", err)
	}
	// rbac
	if !subj.HasRole(domain.RoleUser) {
		return errs.NewBllForbiddenError("UserFacadeImpl.ChangePassword", "user is not a user", nil)
	}
	// logic
	err = uf.changePasswordUC.ChangePassword(ctx, subj.ID, changePassword.OldPassword, changePassword.NewPassword)
	if err != nil {
		return err
	}

	return nil
}

// ChangeKeys wraps asymmetric rotation boundaries, unpacking context indices to generate and yield renewed cryptographic public tokens.
func (uf *UserFacadeImpl) ChangeKeys(ctx context.Context) (*dto.ChangedKeysDTO, error) {
	// subject
	subj, err := uf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserFacadeImpl.ChangeKeys", "retrieve subject", err)
	}
	// rbac
	if !subj.HasRole(domain.RoleUser) {
		return nil, errs.NewBllForbiddenError("UserFacadeImpl.ChangeKeys", "user is not a user", nil)
	}
	// logic
	_, publicKey, err := uf.changeKeysUC.ChangeKeys(ctx, subj.ID)
	if err != nil {
		return nil, err
	}

	return mapper.MapKeysToDTO(publicKey), nil
}
