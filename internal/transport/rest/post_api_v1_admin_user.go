package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	_ "github.com/ElfAstAhe/tiny-auth-service/internal/transport"
)

// postAPIV1AdminUser godoc
// @Summary      Создание пользователя
// @Description  Создаёт нового пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        input  body      UserDTO  true  "Пользователь"
// @Success      201    {object}  UserDTO
// @Failure      400    {object}  ErrorDTO "Кривые данные"
// @Failure      403    {object}  ErrorDTO "В доступе отказано"
// @Failure      409    {object}  ErrorDTO "Уже существует"
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/admin/users [post]
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
