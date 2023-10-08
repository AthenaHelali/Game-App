package matchinghandler

import (
	"game-app/service/authservice"
	"game-app/service/matchingservice"
	"game-app/service/presenceservice"
	"game-app/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSVC       matchingservice.Service
	matchingValidator matchingvalidator.Validator
	presenceSvc       presenceservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service,
	matchingSVC matchingservice.Service, matchingValidator matchingvalidator.Validator,
	presenceSvc presenceservice.Service) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		matchingSVC:       matchingSVC,
		matchingValidator: matchingValidator,
		presenceSvc:       presenceSvc,
	}
}
