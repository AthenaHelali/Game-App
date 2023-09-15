package userhandler

import (
	"game-app/service/authservice"
	"game-app/service/user"
	"game-app/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       user.Service
	userValidator uservalidator.Validator
}

func New(authConfig authservice.Config, authSvc authservice.Service, userSvc user.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authConfig, authSvc, userSvc, userValidator,
	}

}
