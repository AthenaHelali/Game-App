package matchinghandler

import (
	_ "game-app/docs"
	"game-app/internal/param"
	"game-app/internal/pkg/claim"
	"game-app/internal/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

// addToWaitingList godoc
// @Summary Add a user to the waiting list
// @Description Add a user to the waiting list for a matching service.
// @Tags waiting-list
// @Accept json
// @Produce json
// @Param request body param.AddToWaitingListRequest true "Request body containing user information"
// @Security ApiKeyAuth
// @Success 200 {object} param.AddToWaitingListResponse
// @Failure 400 {object} error
// @Failure 401 {object} error
// @Failure 403 {object} error
// @Failure 500 {object} error
// @Router /matching/add-to-waiting-list [post]
func (h Handler) addToWaitingList(c echo.Context) error {
	var req param.AddToWaitingListRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	claims := claim.GetClaimFromEchoContext(c)
	req.UserID = claims.UserID
	if err := h.matchingValidator.ValidateAddToWaitingListRequest(req); err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	resp, err := h.matchingSVC.AddToWaitingList(req)
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)

}
