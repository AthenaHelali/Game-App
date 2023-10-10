package userhandler

import (
	"game-app/internal/service/authservice"
	"game-app/internal/service/presenceservice"
	"game-app/internal/service/user"
	"game-app/internal/validator/uservalidator"
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
