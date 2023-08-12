package httpmsg

import (
	"game-app/pkg/richerror"
	"net/http"
)

func HTTPCodeAndMessage(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.GetMessage()
		code := kindToHTTPStatuseCode(re.GetKind())
		if code >= 500 {
			msg = "internal server error"
		}
		return msg, code
	default:
		return err.Error(), http.StatusBadRequest

	}
}
func kindToHTTPStatuseCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	default:
		return http.StatusBadRequest

	}
}
