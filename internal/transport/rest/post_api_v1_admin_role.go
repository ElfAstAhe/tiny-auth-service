package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// postAPIV1AdminRole godoc
// @Summary      Создание роли
// @Description  Сохраняет новые тестовые данные
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        input  body      RoleDTO  true  "Роль"
// @Success      201    {object}  RoleDTO
// @Failure      400    {object}  ErrorDTO "Кривые данные"
// @Failure      403    {object}  ErrorDTO "В доступе отказано"
// @Failure      409    {object}  ErrorDTO "Уже существует"
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/roles [post]
func (cr *AppChiRouter) postAPIV1AdminRole(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("postAPIV1AdminRole start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("postAPIV1AdminRole finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.RoleDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.roleAdminFacade.Create(r.Context(), income)
	if err != nil {
		cr.log.Errorf("postAPIV1AdminRole post role error, [%v]", err)

		cr.renderError(rw, err)

		return
	}
	location := r.URL.JoinPath(res.ID)
	rw.Header().Set("Location", location.String())

	cr.renderJSON(rw, http.StatusCreated, res)
}
