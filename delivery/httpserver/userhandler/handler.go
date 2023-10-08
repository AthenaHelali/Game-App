package userhandler

import (
	"game-app/service/authservice"
	"game-app/service/presenceservice"
	"game-app/service/user"
	"game-app/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       user.Service
	userValidator uservalidator.Validator
	presenceSvc   presenceservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service, userSvc user.Service, userValidator uservalidator.Validator, presenceSvc presenceservice.Service) Handler {
	return Handler{
		authConfig:    authConfig,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
		presenceSvc:   presenceSvc,
	}

}
