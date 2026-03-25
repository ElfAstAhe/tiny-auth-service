package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// deleteAPIV1AdminRole godoc
// @Summary      Удаление роли
// @Description  Удаляет запись по её ID
// @Tags         role
// @Param        id   path      string  true  "ID записи" format(string)
// @Success      204  "Запись успешно удалена, тело ответа отсутствует"
// @Failure      403  {object   ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/roles/{id} [delete]
func (cr *AppChiRouter) deleteAPIV1AdminRole(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cr.log.Debugf("deleteAPIV1AdminRole start, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)
	defer cr.log.Debugf("deleteAPIV1AdminRole finish, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)

	err := cr.roleAdminFacade.Delete(r.Context(), id)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderEmpty(rw, http.StatusNoContent)
}
