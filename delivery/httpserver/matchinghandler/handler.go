package matchinghandler

import (
	"game-app/service/authservice"
	"game-app/service/matchingservice"
	"game-app/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSVC       matchingservice.Service
	matchingValidator matchingvalidator.Validator
}

func New(authConfig authservice.Config, authSvc authservice.Service, matchingSVC matchingservice.Service, matchingValidator matchingvalidator.Validator) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		matchingSVC:       matchingSVC,
		matchingValidator: matchingValidator,
	}
}
