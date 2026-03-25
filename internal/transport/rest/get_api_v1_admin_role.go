package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// getAPIV1AdminRole godoc
// @Summary      Получить
// @Description  Получает запись по её ID
// @Tags         test
// @Produce      json
// @Param        id   path      string  true  "ID записи" format(string)
// @Success      200  {object}  RoleDTO "Роль"
// @Failure      403  {object   ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/roles/{id} [get]
func (cr *AppChiRouter) getAPIV1AdminRole(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cr.log.Debugf("getAPIV1AdminRole start, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)
	defer cr.log.Debugf("getAPIV1AdminRole finish, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)

	res, err := cr.roleAdminFacade.Get(r.Context(), id)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
