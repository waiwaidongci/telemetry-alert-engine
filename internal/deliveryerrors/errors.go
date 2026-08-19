package deliveryerrors

import "errors"

var ErrRouteMissing = errors.New("delivery route missing")

type LookupFailure struct {
	RouteID string
	Err     error
}

func (e LookupFailure) Error() string {
	return "route lookup failed: " + e.Err.Error()
}

func (e LookupFailure) Unwrap() error {
	return nil
}
