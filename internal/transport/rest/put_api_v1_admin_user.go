package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// putAPIV1AdminUser godoc
// @Summary      Изменяет пользователя
// @Description  Изменение атрибутов пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id     path      string   true  "ID записи" format(string)
// @Param        input  body      UserDTO  true  "Роль"
// @Success      200    {object}  UserDTO
// @Failure      400    {object}  ErrorDTO
// @Failure      404    {object}  ErrorDTO
// @Failure      409    {object}  ErrorDTO
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/users/{id} [put]
func (cr *AppChiRouter) putAPIV1AdminUser(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cr.log.Debugf("putAPIV1AdminUser start, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)
	defer cr.log.Debugf("putAPIV1AdminUser finish, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)

	var income = &dto.UserDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.userAdminFacade.Change(r.Context(), id, income)
	if err != nil {
		cr.log.Errorf("putAPIV1AdminUser put user error, [%v]", err)

		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
