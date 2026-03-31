package rest

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// getAPIV1AdminRoleSearch godoc
// @Summary      Получить
// @Description  Поиск записи
// @Tags         role
// @Produce      json
// @Param        name   query      string  true  "name записи" format(string)
// @Success      200  {object}  RoleDTO "Роль"
// @Failure      403  {object}  ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/roles/search [get]
func (cr *AppChiRouter) getAPIV1AdminRoleSearch(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("getAPIV1AdminRoleSearch start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("getAPIV1AdminRoleSearch finish, requestID [%s]", middleware.GetReqID(r.Context()))

	name := cr.getQueryString(r, "name", "")
	if name == "" {
		cr.renderError(rw, errs.NewInvalidArgumentError("name", ""))

		return
	}

	res, err := cr.roleAdminFacade.GetByName(r.Context(), name)
	if err != nil {
		cr.log.Errorf("getAPIV1AdminRoleSearch get role error, [%v]", err)

		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
