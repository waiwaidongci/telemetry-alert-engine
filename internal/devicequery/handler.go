package devicequery

func HTTPStatus(kind Kind) int {
	if kind == NotFound {
		return 404
	}
	return 500
}
