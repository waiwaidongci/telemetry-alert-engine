package notificationerrors

import "net/http"

func HTTPStatus(kind Kind) int {
	statuses := map[Kind]int{
		KindMissing: http.StatusInternalServerError,
		KindSystem:  http.StatusInternalServerError,
	}
	switch kind {
	case KindMissing:
		return statuses[KindMissing]
	default:
		return statuses[KindSystem]
	}
}
