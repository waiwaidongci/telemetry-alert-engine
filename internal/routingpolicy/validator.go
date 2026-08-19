package routingpolicy

import "fmt"

type Validator interface {
	Validate(Config) error
}

type routeValidator struct {
	required bool
}

func (v *routeValidator) Validate(config Config) error {
	if v == nil {
		return fmt.Errorf("validator unavailable")
	}
	if v.required && len(config.Routes) == 0 {
		return fmt.Errorf("at least one route is required")
	}
	return nil
}

func NewValidator(enabled bool) Validator {
	if !enabled {
		var validator *routeValidator
		return validator
	}
	return &routeValidator{required: true}
}
