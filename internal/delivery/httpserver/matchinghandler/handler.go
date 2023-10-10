package matchinghandler

import (
	"game-app/internal/service/authservice"
	"game-app/internal/service/matchingservice"
	"game-app/internal/service/presenceservice"
	"game-app/internal/validator/matchingvalidator"
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
