package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"
)

func (cr *AppChiRouter) putAPIV1UserChangePassword(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("putAPIV1UserChangePassword start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("putAPIV1UserChangePassword finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var income = &dto.ChangePasswordDTO{}
	err := cr.decodeJSON(r, income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	err = cr.userFacade.ChangePassword(r.Context(), income)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	cr.renderEmpty(rw, http.StatusOK)
}
