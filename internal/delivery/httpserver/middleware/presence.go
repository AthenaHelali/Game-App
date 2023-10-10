package middleware

import (
	"game-app/internal/param"
	"game-app/internal/pkg/claim"
	"game-app/internal/pkg/errormessage"
	"game-app/internal/pkg/timestamp"
	"game-app/internal/service/presenceservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			claims := claim.GetClaimFromEchoContext(c)
			_, err = service.Upsert(c.Request().Context(), param.UpsertPresenceRequest{
				UserID:    claims.UserID,
				Timestamp: timestamp.Now(),
			})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errormessage.ErrorMsgSomeThingWentWrong,
				})

			}
			return next(c)
		}
	}
}
