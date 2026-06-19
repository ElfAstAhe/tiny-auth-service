package rest

import (
	"net/http"

	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport"
	"github.com/go-chi/chi/v5/middleware"
)

// getAPIV1AdminRoles godoc
// @Summary      Получить
// @Description  Получить список
// @Tags         role
// @Produce      json
// @Param        limit   query   int  false  "limit row count, max 1000" format(int)
// @Param        offset  query   int  false  "offset, min 0, max n" format(int)
// @Success      200  {array}  RoleDTO "Набор ролей"
// @Failure      400  {object} ErrorDTO
// @Failure      403  {object} ErrorDTO "В доступе отказано"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/roles [get]
func (cr *AppChiRouter) getAPIV1AdminRoles(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("getAPIV1AdminRoles start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("getAPIV1AdminRoles finish, requestID [%s]", middleware.GetReqID(r.Context()))

	limit := libhttp.GetQueryIntDefault(r, "limit", transport.DefaultListLimit)
	offset := libhttp.GetQueryIntDefault(r, "offset", transport.DefaultListOffset)

	res, err := cr.roleAdminFacade.List(r.Context(), limit, offset)
	if err != nil {
		cr.log.Errorf("getAPIV1AdminRoles list roles error, [%v]", err)

		libhttp.RenderError(rw, err, mapToHTTPStatus)

		return
	}

	libhttp.RenderJSON(rw, http.StatusOK, res, mapToHTTPStatus)
}
