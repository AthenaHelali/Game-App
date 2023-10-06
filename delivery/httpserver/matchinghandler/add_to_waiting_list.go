package matchinghandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"game-app/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
