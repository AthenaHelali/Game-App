package httpserver

import (
	_ "game-app/docs"
	"github.com/labstack/echo/v4"
	"net/http"
)

// healthCheck godoc
// @Summary Health Check
// @Description Check the health status of the server.
// @ID health-check
// @Produce json
// @Success 200 {object} error
// @Router /health-check [get]
func (s Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good",
	})

}
