package matchinghandler

import (
	middleware2 "game-app/internal/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetMatchingRoutes(e *echo.Echo) {
	userGroup := e.Group("/matching")

	userGroup.POST("/add-to-waiting-list", h.addToWaitingList, middleware2.Auth(h.authSvc, h.authConfig), middleware2.UpsertPresence(h.presenceSvc))

}
