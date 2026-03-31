package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// getAPIV1UserProfile godoc
// @Summary      Получить
// @Description  Получает запись по её ID
// @Tags         profile
// @Produce      json
// @Success      200  {object}  ProfileDTO "Профиль"
// @Failure      403  {object   ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/users/profile [get]
func (cr *AppChiRouter) getAPIV1UserProfile(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("getAPIV1UserProfile start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("getAPIV1UserProfile finish, requestID [%s]", middleware.GetReqID(r.Context()))

	res, err := cr.userFacade.Profile(r.Context())
	if err != nil {
		cr.log.Errorf("getAPIV1UserProfile err: [%v]", err)

		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
