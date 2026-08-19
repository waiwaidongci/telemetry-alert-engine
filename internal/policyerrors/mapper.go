package policyerrors

func HTTPStatus(err error) int {
	if err == nil {
		return 200
	}
	return 500
}

func Retryable(err error) bool {
	if err == nil {
		return false
	}
	return true
}
