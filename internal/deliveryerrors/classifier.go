package deliveryerrors

import "errors"

func Classify(err error) int {
	switch {
	case err == nil:
		return 200
	case errors.Is(err, ErrRouteMissing):
		return 404
	default:
		return 500
	}
}

func Permanent(err error) bool {
	return errors.Is(err, ErrRouteMissing)
}
