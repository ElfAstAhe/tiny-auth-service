package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// putAPIV1UserChangePassword godoc
// @Summary      Смена пароля пользователя
// @Description  Изменяет пароль пользователя
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        input  body      ChangePasswordDTO  true  "Смена пароля"
// @Success      200    "Успех (пустое тело)"
// @Failure      400    {object}  ErrorDTO
// @Failure      404    {object}  ErrorDTO
// @Failure      409    {object}  ErrorDTO
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/users/password [put]
func (cr *AppChiRouter) putAPIV1UserChangePassword(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("putAPIV1UserChangePassword start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("putAPIV1UserChangePassword finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.ChangePasswordDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	err = cr.userFacade.ChangePassword(r.Context(), income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderEmpty(rw, http.StatusOK)
}
