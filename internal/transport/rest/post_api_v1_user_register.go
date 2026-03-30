package rest

import (
	"net/http"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/go-chi/chi/v5/middleware"
)

// postAPIV1UserRegister godoc
// @Summary      Регистрация
// @Description  Регистрация пользователя
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        input  body      RegisterDTO  true  "register информация"
// @Success      200    {object}  ProfileDTO
// @Failure      400    {object}  ErrorDTO "Кривые данные"
// @Failure      409    {object}  ErrorDTO "Уже существует"
// @Failure      500    "Внутренняя ошибка сервера (пустое тело)"
// @Router       /api/v1/user/register [post]
func (cr *AppChiRouter) postAPIV1UserRegister(rw http.ResponseWriter, r *http.Request) {
	cr.log.Debugf("postAPIV1UserRegister start, requestID [%s]", middleware.GetReqID(r.Context()))
	defer cr.log.Debugf("postAPIV1UserRegister finish, requestID [%s]", middleware.GetReqID(r.Context()))

	var incomeDTO = &dto.RegisterDTO{}
	err := cr.decodeJSON(r, incomeDTO)
	if err != nil {
		cr.renderError(rw, err)

		return
	}

	res, err := cr.userFacade.Register(r.Context(), incomeDTO)
	if err != nil {
		cr.log.Errorf("userFacade register error: [%v]", err)

		cr.renderError(rw, err)

		return
	}

	cr.renderJSON(rw, http.StatusOK, res)
}
