package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (cr *AppChiRouter) getAPIV1UserProfile(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("getAPIV1UserProfile start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("getAPIV1UserProfile finish, requestID [%s]", middleware.GetReqID(r.Context()))

	res, err := cr.userFacade.Profile(r.Context())
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
