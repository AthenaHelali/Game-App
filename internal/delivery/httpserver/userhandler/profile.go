package userhandler

import (
	_ "game-app/docs"
	"game-app/internal/param"
	"game-app/internal/pkg/claim"
	"game-app/internal/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

// userProfile godoc
// @Summary Get user profile
// @Description Get the user's profile information.
// @ID getUserProfile
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header (Bearer token)"
// @Success 200 {object} param.ProfileResponse "User profile retrieved successfully"
// @Failure 400 {object} error "Bad request"
// @Failure 401 {object} error "Unauthorized"
// @Failure 404 {object} error "User not found"
// @Failure 500 {object} error "Internal server error"
// @Router /user/profile [get]
func (h Handler) userProfile(c echo.Context) error {
	cl := claim.GetClaimFromEchoContext(c)

	resp, err := h.userSvc.Profile(c.Request().Context(), param.ProfileRequest{UserID: cl.UserID})
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)

}
