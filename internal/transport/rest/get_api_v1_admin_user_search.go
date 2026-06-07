package rest

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// getAPIV1AdminUserSearch godoc
// @Summary      Получить
// @Description  Поиск записи
// @Tags         user
// @Produce      json
// @Param        name   query      string  true  "name записи" format(string)
// @Success      200  {object}  UserDTO "Пользователь"
// @Failure      403  {object}  ErrorDTO "В доступе отказано"
// @Failure      404  {object}  ErrorDTO "Запись не найдена"
// @Failure      500  "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/users/search [get]
func (cr *AppChiRouter) getAPIV1AdminUserSearch(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("getAPIV1AdminUserSearch start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("getAPIV1AdminUserSearch finish, requestID [%s]", middleware.GetReqID(r.Context()))

	name := libhttp.GetQueryStringDefault(r, "name", "")
	if name == "" {
		libhttp.RenderError(rw, errs.NewInvalidArgumentError("name", ""), mapToHTTPStatus)

		return
	}

	res, err := cr.userAdminFacade.GetByName(r.Context(), name)
	if err != nil {
		cr.log.Errorf("getAPIV1AdminUserSearch get user error, [%v]", err)

		libhttp.RenderError(rw, err, mapToHTTPStatus)

		return
	}

	libhttp.RenderJSON(rw, http.StatusOK, res, mapToHTTPStatus)
}
