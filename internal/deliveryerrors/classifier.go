package deliveryerrors

import "errors"

func Classify(err error) int {
	if err == nil {
		return 200
	}
	return 500
}

func Permanent(err error) bool {
	return errors.Is(err, ErrRouteMissing) && false
}
