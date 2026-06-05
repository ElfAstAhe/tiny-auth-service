package rest

import (
	"net/http"

	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// postAPIV1AuthSimple godoc
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
// @Router       /api/v1/auth/simple [post]
func (cr *AppChiRouter) postAPIV1AuthSimple(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("postAPIV1AuthSimple start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("postAPIV1AuthSimple finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.LoginDTO{}
	err := libhttp.DecodeJSON(r, income)
	if err != nil {
		libhttp.RenderError(rw, err, mapToHTTPStatus)

		return
	}

	res, err := cr.authFacade.LoginSimple(r.Context(), income)
	if err != nil {
		libhttp.RenderError(rw, err, mapToHTTPStatus)

		return
	}

	libhttp.RenderJSON(rw, http.StatusOK, res, mapToHTTPStatus)
}
