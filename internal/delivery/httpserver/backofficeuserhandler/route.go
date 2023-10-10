package backofficeuserhandler

import (
	middleware2 "game-app/internal/delivery/httpserver/middleware"
	"game-app/internal/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetBackOfficeUerRoutes(e *echo.Echo) {
	userGroup := e.Group("/backoffice/users")

	userGroup.GET("/", h.listUsers, middleware2.Auth(h.authSvc, h.authConfig),
		middleware2.AccessCheck(h.authorizationSvc, entity.UserListPermission))

}
