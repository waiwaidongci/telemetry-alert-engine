package policyerrors

import "errors"

func HTTPStatus(err error) int {
	if err == nil {
		return 200
	}
	var limitError *LimitError
	if errors.As(err, &limitError) {
		return 429
	}
	return 500
}

func Retryable(err error) bool {
	if err == nil {
		return false
	}
	var classified interface{ Retryable() bool }
	if errors.As(err, &classified) {
		return classified.Retryable()
	}
	return true
}
