package backofficeuserhandler

import (
	"game-app/delivery/httpserver/middleware"
	"game-app/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetBackOfficeUerRoutes(e *echo.Echo) {
	userGroup := e.Group("/backoffice/users")

	userGroup.GET("/", h.listUsers, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))

}
