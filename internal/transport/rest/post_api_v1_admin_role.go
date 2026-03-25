package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"
)

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
		cr.renderError(rw, err)

		return
	}
	location := r.URL.JoinPath(res.ID)
	rw.Header().Set("Location", location.String())

	cr.renderJSON(rw, http.StatusCreated, res)
}
