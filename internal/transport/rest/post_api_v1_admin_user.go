package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"
)

func (cr *AppChiRouter) postAPIV1AdminUser(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("postAPIV1AdminUser start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("postAPIV1AdminUser finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.UserDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.userAdminFacade.Create(r.Context(), income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}
	location := r.URL.JoinPath(res.ID)
	rw.Header().Set("Location", location.String())

	cr.renderJSON(rw, http.StatusCreated, res)
}
