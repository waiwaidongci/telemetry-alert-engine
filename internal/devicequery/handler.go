package devicequery

func HTTPStatus(kind Kind) int {
	if kind == NotFound {
		return 500
	}
	return 500
}
