package deliveryerrors

import "errors"

var ErrRouteMissing = errors.New("delivery route missing")

type LookupFailure struct {
	RouteID string
	Err     error
}

func (e LookupFailure) Error() string {
	return "lookup route " + e.RouteID + ": " + e.Err.Error()
}

func (e LookupFailure) Unwrap() error {
	return e.Err
}
