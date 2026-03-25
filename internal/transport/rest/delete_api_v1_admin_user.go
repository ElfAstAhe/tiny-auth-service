package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// deleteAPIV1AdminUser godoc
// @Summary      Удаление пользователя
// @Description  Удаляет запись по её ID
// @Tags         User
// @Param        id   path      string  true  "ID записи" format(string)
// @Success      204  "Запись успешно удалена, тело ответа отсутствует"
// @Failure      403  {object   ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/users/{id} [delete]
func (cr *AppChiRouter) deleteAPIV1AdminUser(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cr.log.Debugf("deleteAPIV1AdminUser start, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)
	defer cr.log.Debugf("deleteAPIV1AdminUser finish, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)

	err := cr.userAdminFacade.Delete(r.Context(), id)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderEmpty(rw, http.StatusNoContent)
}
