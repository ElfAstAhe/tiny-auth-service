package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
