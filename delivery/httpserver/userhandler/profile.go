package userhandler

import (
	"game-app/param"
	"game-app/pkg/constant"
	"game-app/pkg/httpmsg"
	"game-app/service/authservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getClaims(c echo.Context) *authservice.Claims {
	return c.Get(constant.AuthMiddlewareContextKey).(*authservice.Claims)
}

func (h Handler) userProfile(c echo.Context) error {

	cl := getClaims(c)

	resp, err := h.userSvc.Profile(param.ProfileRequest{UserID: cl.UserID})
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)

}
