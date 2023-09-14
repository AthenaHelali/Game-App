package userhandler

import (
	"game-app/service/authservice"
	"game-app/service/user"
	"game-app/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       user.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc user.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc, userSvc, userValidator,
	}

}
