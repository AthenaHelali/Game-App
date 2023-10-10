package httpserver

import (
	"fmt"
	"game-app/internal/config"
	"game-app/internal/delivery/httpserver/backofficeuserhandler"
	"game-app/internal/delivery/httpserver/matchinghandler"
	"game-app/internal/delivery/httpserver/userhandler"
	"game-app/internal/service/authorizationservice"
	"game-app/internal/service/authservice"
	"game-app/internal/service/backofficeuserservice"
	"game-app/internal/service/matchingservice"
	"game-app/internal/service/presenceservice"
	"game-app/internal/service/user"
	"game-app/internal/validator/matchingvalidator"
	"game-app/internal/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"log"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
	Router                *echo.Echo
}

func New(config config.Config, authSvc authservice.Service, userSvc user.Service, backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service,
	userValidator uservalidator.Validator, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator, presenceSvc presenceservice.Service) Server {
	return Server{
		Router:                echo.New(),
		config:                config,
		userHandler:           userhandler.New(config.Auth, authSvc, userSvc, userValidator, presenceSvc),
		backofficeUserHandler: backofficeuserhandler.New(config.Auth, authSvc, backofficeUserSvc, authorizationSvc),
		matchingHandler:       matchinghandler.New(config.Auth, authSvc, matchingSvc, matchingValidator, presenceSvc),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	s.Router.GET("/health-check", s.healthCheck)

	s.userHandler.SetUerRoutes(s.Router)
	s.backofficeUserHandler.SetBackOfficeUerRoutes(s.Router)
	s.matchingHandler.SetMatchingRoutes(s.Router)

	s.Router.GET("/swagger/*", echoSwagger.WrapHandler)

	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	log.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		log.Println("router start error", err)
	}
}
