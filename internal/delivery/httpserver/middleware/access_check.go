package middleware

import (
	"game-app/internal/entity"
	"game-app/internal/pkg/claim"
	"game-app/internal/pkg/errormessage"
	"game-app/internal/service/authorizationservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AccessCheck(service authorizationservice.Service, permissions ...entity.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			claims := claim.GetClaimFromEchoContext(c)
			isAllowed, err := service.CheckAccess(c.Request().Context(), claims.UserID, claims.Role, permissions...)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errormessage.ErrorMsgSomeThingWentWrong,
				})
			}
			if !isAllowed {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": errormessage.ErrorMsgUserNotAllowed,
				})
			}
			return next(c)
		}
	}
}
