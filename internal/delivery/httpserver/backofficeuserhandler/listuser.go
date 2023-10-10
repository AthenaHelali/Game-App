package backofficeuserhandler

import (
	_ "game-app/docs"
	"game-app/internal/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

// listUsers godoc
// @Summary List all users
// @Description Retrieve a list of all users from the back office.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} entity.User
// @Failure 400 {object} error
// @Failure 401 {object} error
// @Failure 403 {object} error
// @Failure 500 {object} error
// @Router /backoffice/users/ [get]
func (h Handler) listUsers(c echo.Context) error {
	list, err := h.backofficeUserSvc.ListAllUsers()
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": list,
	})
}
