package ruleconfig

type Validator interface{ Validate(string) error }
type ruleValidator struct{}

func (*ruleValidator) Validate(string) error { return nil }
func DisabledValidator() Validator           { return nil }
