package claim

import (
	cfg "game-app/internal/config"
	"game-app/internal/service/authservice"
	"github.com/labstack/echo/v4"
)

func GetClaimFromEchoContext(c echo.Context) *authservice.Claims {
	return c.Get(cfg.AuthMiddlewareContextKey).(*authservice.Claims)
}
