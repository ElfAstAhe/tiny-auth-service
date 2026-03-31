package rest

import (
	"net/http"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
	"github.com/go-chi/chi/v5/middleware"
)

// getAPIV1AdminUsers godoc
// @Summary      Получить
// @Description  Получить список
// @Tags         user
// @Produce      json
// @Param        limit   query   int  false  "limit row count, max 1000" format(int)
// @Param        offset  query   int  false  "offset, min 0, max n" format(int)
// @Success      200  {array}  UserDTO "Набор пользователей"
// @Failure      400  {object} ErrorDTO
// @Failure      403  {object} ErrorDTO "В доступе отказано"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/users [get]
func (cr *AppChiRouter) getAPIV1AdminUsers(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("getAPIV1AdminUsers start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("getAPIV1AdminUsers finish, requestID [%s]", middleware.GetReqID(r.Context()))

	limit, err := cr.getQueryInt(r, "limit", transport.DefaultListLimit)
	if err != nil {
		cr.renderError(rw, err)

		return
	}
	offset, err := cr.getQueryInt(r, "offset", transport.DefaultListOffset)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.userAdminFacade.List(r.Context(), limit, offset)
	if err != nil {
		cr.log.Errorf("getAPIV1AdminUsers list users error, [%v]", err)

		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
