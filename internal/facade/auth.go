package facade

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/mapper"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

// AuthFacade defines the core authentication orchestration contract managing token generation and session validation workflows.
type AuthFacade interface {
	// Login processes full asymmetric cryptographic identity validations and yields active token structures.
	Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error)
	// LoginSimple handles lightweight authentication bypassing intensive key pair operations.
	LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error)
}

// AuthFacadeImpl structures sessions boundary routers, conducting metadata transformations and payload assertions.
type AuthFacadeImpl struct {
	jwtHelper     *helper.JWTHelper          // Framework core security utility managing serialization of token cryptographic layouts
	loginUC       usecase.LoginUseCase       // Downstream usecase executing public full cryptographic authentication scenarios
	loginSimpleUC usecase.LoginSimpleUseCase // Downstream usecase handling simplified machine session credentials processing
}

// Compile-time interface compliance verification
var _ AuthFacade = (*AuthFacadeImpl)(nil)

// NewAuthFacade acts as a factory constructor embedding modular usecase components and token serialization utilities.
func NewAuthFacade(jwtHelper *helper.JWTHelper, loginUC usecase.LoginUseCase, loginSimpleUC usecase.LoginSimpleUseCase) *AuthFacadeImpl {
	return &AuthFacadeImpl{
		jwtHelper:     jwtHelper,
		loginUC:       loginUC,
		loginSimpleUC: loginSimpleUC,
	}
}

// Login validates inbound credentials boundaries, delegates state processing, and structures fully marshaled secure tokens strings.
func (af *AuthFacadeImpl) Login(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	if err := af.validate(login); err != nil {
		return nil, errs.NewBllValidateError("AuthFacadeImpl.Login", "validate income failed", err)
	}

	token, _, err := af.loginUC.Login(ctx, login.Username, login.Password)
	if err != nil {
		return nil, errs.NewBllUnauthorizedError("AuthFacadeImpl.Login", "unauthorized", err)
	}
	tokenStr, err := af.jwtHelper.BuildTokenStr(token)
	if err != nil {
		return nil, errs.NewBllUnauthorizedError("AuthFacadeImpl.Login", "unauthorized", err)
	}
	// refreshTokenStr, err := af.jwtHelper.BuildTokenStr(refreshToken)
	refreshTokenStr := ""
	//if err != nil {
	//	return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.Login", "unauthorized", err)
	//}

	return mapper.MapAuthToDTO(tokenStr, refreshTokenStr), nil
}

// LoginSimple processes lightweight credentials matching, constructing simplified fast-token wire contracts.
func (af *AuthFacadeImpl) LoginSimple(ctx context.Context, login *dto.LoginDTO) (*dto.LoggedInDTO, error) {
	if err := af.validate(login); err != nil {
		return nil, errs.NewBllValidateError("AuthFacadeImpl.LoginSimple", "validate income failed", err)
	}

	token, _, err := af.loginSimpleUC.Login(ctx, login.Username, login.Password)
	if err != nil {
		return nil, errs.NewBllUnauthorizedError("AuthFacadeImpl.LoginSimple", "unauthorized", err)
	}
	tokenStr, err := af.jwtHelper.BuildTokenStr(token)
	if err != nil {
		return nil, errs.NewBllUnauthorizedError("AuthFacadeImpl.LoginSimple", "unauthorized", err)
	}
	//refreshTokenStr, err := af.jwtHelper.BuildTokenStr(refreshToken)
	//if err != nil {
	//	return nil, domerrs.NewBllUnauthorizedError("AuthFacadeImpl.LoginSimple", "unauthorized", err)
	//}
	refreshTokenStr := ""

	return mapper.MapAuthToDTO(tokenStr, refreshTokenStr), nil
}

// validate executes semantic checking across required inbound object attributes prior to starting execution loops.
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
