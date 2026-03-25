package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"
)

func (cr *AppChiRouter) postAPIV1Auth(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("postAPIV1Auth start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("postAPIV1Auth finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.LoginDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.authFacade.Login(r.Context(), income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
