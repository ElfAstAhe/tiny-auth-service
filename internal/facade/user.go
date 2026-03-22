package facade

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

type UserFacade interface {
	Profile(ctx context.Context) (*dto.ProfileDTO, error)
	ChangePassword(ctx context.Context, changePassword *dto.ChangePasswordDTO) error
	ChangeKeys(ctx context.Context) (*dto.ChangedKeysDTO, error)
}

type UserFacadeImpl struct {
	authHelper       *auth.Helper
	profileUC        usecase.ProfileUseCase
	changePasswordUC usecase.ChangePasswordUseCase
	changeKeysUC     usecase.ChangeKeysUseCase
}

func NewUserFacade(authHelper *auth.Helper, profileUC usecase.ProfileUseCase, changePasswordUC usecase.ChangePasswordUseCase, changeKeysUC usecase.ChangeKeysUseCase) *UserFacadeImpl {
	return &UserFacadeImpl{
		authHelper:       authHelper,
		profileUC:        profileUC,
		changePasswordUC: changePasswordUC,
		changeKeysUC:     changeKeysUC,
	}
}

func (uf *UserFacadeImpl) Profile(ctx context.Context) (*dto.ProfileDTO, error) {
	subj, err := uf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserFacadeImpl.Profile", "retrieve subject", err)
	}
	res, err := uf.profileUC.Get(ctx, subj.Name)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserModelToProfileDTO(res), nil
}

func (uf *UserFacadeImpl) ChangePassword(ctx context.Context, changePassword *dto.ChangePasswordDTO) error {
	subj, err := uf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return errs.NewBllForbiddenError("UserFacadeImpl.ChangePassword", "retrieve subject", err)
	}

	err = uf.changePasswordUC.ChangePassword(ctx, subj.ID, changePassword.OldPassword, changePassword.NewPassword)
	if err != nil {
		return err
	}

	return nil
}

func (uf *UserFacadeImpl) ChangeKeys(ctx context.Context) (*dto.ChangedKeysDTO, error) {
	subj, err := uf.authHelper.SubjectFromContext(ctx)
	if err != nil {
		return nil, errs.NewBllForbiddenError("UserFacadeImpl.ChangeKeys", "retrieve subject", err)
	}

	_, publicKey, err := uf.changeKeysUC.ChangeKeys(ctx, subj.ID)
	if err != nil {
		return nil, err
	}

	return mapper.MapKeysToDTO(publicKey), nil
}
