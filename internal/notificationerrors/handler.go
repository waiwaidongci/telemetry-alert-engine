package notificationerrors

import "net/http"

func HTTPStatus(kind Kind) int {
	switch kind {
	case KindMissing:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
