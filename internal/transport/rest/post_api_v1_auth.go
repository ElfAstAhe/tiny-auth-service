package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// postAPIV1Auth godoc
// @Summary      Аутентификация
// @Description  Аутентификация и авторизация пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      LoginDTO  true  "login информация"
// @Success      200    {object}  LoggedInDTO
// @Failure      400    {object}  ErrorDTO "Кривые данные"
// @Failure      401    {object}  ErrorDTO "Не авторизован"
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/auth [post]
func (cr *AppChiRouter) postAPIV1Auth(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("postAPIV1Auth start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("postAPIV1Auth finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.LoginDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.authFacade.Login(r.Context(), income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
