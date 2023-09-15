package httpserver

import (
	"fmt"
	"game-app/config"
	"game-app/delivery/httpserver/userhandler"
	"game-app/service/authservice"
	"game-app/service/user"
	"game-app/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc user.Service, userValidator uservalidator.Validator) Server {
	return Server{
		config:      config,
		userHandler: userhandler.New(config.Auth, authSvc, userSvc, userValidator),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)
	s.userHandler.SetUserRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
