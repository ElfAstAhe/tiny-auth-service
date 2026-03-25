package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (cr *AppChiRouter) putAPIV1UserChangeKeys(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("putAPIV1UserChangeKeys start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("putAPIV1UserChangeKeys finish, requestID [%s]", middleware.GetReqID(r.Context()))

	res, err := cr.userFacade.ChangeKeys(r.Context())
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
