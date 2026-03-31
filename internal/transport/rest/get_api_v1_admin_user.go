package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// getAPIV1AdminUser godoc
// @Summary      Получить
// @Description  Получает запись по её ID
// @Tags         user
// @Produce      json
// @Param        id   path      string  true  "ID записи" format(string)
// @Success      200  {object}  UserDTO "Пользователь"
// @Failure      403  {object   ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/users/{id} [get]
func (cr *AppChiRouter) getAPIV1AdminUser(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cr.log.Debugf("getAPIV1AdminUser start, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)
	defer cr.log.Debugf("getAPIV1AdminUser finish, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)

	res, err := cr.userAdminFacade.Get(r.Context(), id)
	if err != nil {
		cr.log.Errorf("getAPIV1AdminUser get user error, [%v]", err)

		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
