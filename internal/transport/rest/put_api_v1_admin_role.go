package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// putAPIV1AdminRole godoc
// @Summary      Изменяет роль
// @Description  Изменение атрибутов роли
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id     path      string   true  "ID записи" format(string)
// @Param        input  body      RoleDTO  true  "Роль"
// @Success      200    {object}  RoleDTO
// @Failure      400    {object}  ErrorDTO
// @Failure      404    {object}  ErrorDTO
// @Failure      409    {object}  ErrorDTO
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/roles/{id} [put]
func (cr *AppChiRouter) putAPIV1AdminRole(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cr.log.Debugf("putAPIV1AdminRole start, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)
	defer cr.log.Debugf("putAPIV1AdminRole finish, requestID [%s] path param [%s]", middleware.GetReqID(r.Context()), id)

	var income = &dto.RoleDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.roleAdminFacade.Change(r.Context(), id, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
