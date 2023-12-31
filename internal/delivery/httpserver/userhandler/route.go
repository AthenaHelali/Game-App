package userhandler

import (
	middleware2 "game-app/internal/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUerRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)

	userGroup.POST("/login", h.userLogin)

	userGroup.GET("/profile", h.userProfile, middleware2.Auth(h.authSvc, h.authConfig), middleware2.UpsertPresence(h.presenceSvc))

}
