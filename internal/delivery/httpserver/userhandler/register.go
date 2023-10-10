package userhandler

import (
	_ "game-app/docs"
	"game-app/internal/param"
	"game-app/internal/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

// userRegister godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param body body param.RegisterRequest true "User registration request"
// @Success 201 {object} param.RegisterResponse
// @Failure 400 {object} error
// @Router /api/users/register [post]
func (h Handler) userRegister(c echo.Context) error {
	var uReq param.RegisterRequest

	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}
	if err := h.userValidator.ValidateRegisterRequest(uReq); err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	response, err := h.userSvc.Register(c.Request().Context(), uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}
