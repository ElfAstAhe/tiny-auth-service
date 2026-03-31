package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// putAPIV1UserChangeKeys godoc
// @Summary      Изменяет роль
// @Description  Изменение атрибутов роли
// @Tags         profile
// @Produce      json
// @Success      200    {object}  ChangedKeysDTO
// @Failure      400    {object}  ErrorDTO
// @Failure      404    {object}  ErrorDTO
// @Failure      409    {object}  ErrorDTO
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/users/keys [put]
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
