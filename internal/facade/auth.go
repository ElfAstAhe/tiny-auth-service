package facade

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	domerrs "github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

type AuthFacade interface {
	Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error)
	LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error)
}

type AuthFacadeImpl struct {
	jwtHelper     *helper.JWTHelper
	loginUC       usecase.LoginUseCase
	loginSimpleUC usecase.LoginSimpleUseCase
}

func NewAuthFacade(jwtHelper *helper.JWTHelper, loginUC usecase.LoginUseCase, loginSimpleUC usecase.LoginSimpleUseCase) *AuthFacadeImpl {
	return &AuthFacadeImpl{
		jwtHelper:     jwtHelper,
		loginUC:       loginUC,
		loginSimpleUC: loginSimpleUC,
	}
}

func (af *AuthFacadeImpl) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	if err := af.validate(login); err != nil {
		return nil, domerrs.NewBllValidateError("AuthFacadeImpl.Login", "validate income failed", err)
	}

	token, refreshToken, err := af.loginUC.Login(ctx, login.Username, login.Password)
	if err != nil {
		return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.Login", "unauthorized", err)
	}
	tokenStr, err := af.jwtHelper.BuildTokenStr(token)
	if err != nil {
		return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.Login", "unauthorized", err)
	}
	refreshTokenStr, err := af.jwtHelper.BuildTokenStr(refreshToken)
	if err != nil {
		return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.Login", "unauthorized", err)
	}

	return mapper.MapAuthToDTO(tokenStr, refreshTokenStr), nil
}

func (af *AuthFacadeImpl) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	if err := af.validate(login); err != nil {
		return nil, domerrs.NewBllValidateError("AuthFacadeImpl.LoginSimple", "validate income failed", err)
	}

	token, refreshToken, err := af.loginSimpleUC.Login(ctx, login.Username, login.Password)
	if err != nil {
		return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.LoginSimple", "unauthorized", err)
	}
	tokenStr, err := af.jwtHelper.BuildTokenStr(token)
	if err != nil {
		return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.LoginSimple", "unauthorized", err)
	}
	refreshTokenStr, err := af.jwtHelper.BuildTokenStr(refreshToken)
	if err != nil {
		return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.LoginSimple", "unauthorized", err)
	}

	return mapper.MapAuthToDTO(tokenStr, refreshTokenStr), nil
}

func (af *AuthFacadeImpl) validate(loginDTO *dto.LoginDTO) error {
	if loginDTO == nil {
		return errs.NewInvalidArgumentError("loginDTO", "must not be nil")
	}
	if loginDTO.Username == "" {
		return errs.NewInvalidArgumentError("loginDTO.username", "must not be empty")
	}
	if loginDTO.Password == "" {
		return errs.NewInvalidArgumentError("loginDTO.password", "must not be empty")
	}

	return nil
}
