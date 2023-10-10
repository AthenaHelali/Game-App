package backofficeuserhandler

import (
	"game-app/internal/service/authorizationservice"
	"game-app/internal/service/authservice"
	"game-app/internal/service/backofficeuserservice"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	backofficeUserSvc backofficeuserservice.Service
	authorizationSvc  authorizationservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service, backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service) Handler {
	return Handler{
		authConfig, authSvc, backofficeUserSvc, authorizationSvc,
	}

}
