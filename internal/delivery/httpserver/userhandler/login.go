package userhandler

import (
	_ "game-app/docs"
	"game-app/internal/param"
	"game-app/internal/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

// userLogin godoc
// @Summary Login a user
// @Description Login a user with the provided credentials
// @Tags users
// @Accept json
// @Produce json
// @Param body body param.LoginRequest true "User login request"
// @Success 200 {object} param.LoginResponse
// @Failure 400 {object} error
// @Router /api/users/login [post]
func (h Handler) userLogin(c echo.Context) error {
	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	if err := h.userValidator.ValidateLoginRequest(req); err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	response, err := h.userSvc.Login(c.Request().Context(), req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
