package httpserver

import (
	"game-app/service/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userRegister(c echo.Context) error {
	var uReq user.RegisterRequest

	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	response, err := s.userSvc.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

func (s Server) userLogin(c echo.Context) error {
	var req user.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	response, err := s.userSvc.Login(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}

func (s Server) userProfile(c echo.Context) error {

	authToken := c.Request().Header.Get("Authorization")
	claims, err := s.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.Profile(user.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)

}
