package claim

import (
	cfg "game-app/config"
	"game-app/service/authservice"
	"github.com/labstack/echo/v4"
)

func GetClaimFromEchoContext(c echo.Context) *authservice.Claims {
	return c.Get(cfg.AuthMiddlewareContextKey).(*authservice.Claims)
}
