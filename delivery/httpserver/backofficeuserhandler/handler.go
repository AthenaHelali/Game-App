package backofficeuserhandler

import (
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backofficeuserservice"
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
