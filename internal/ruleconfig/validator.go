package ruleconfig

type Validator interface{ Validate(string) error }
type ruleValidator struct{}

func (*ruleValidator) Validate(string) error { return nil }
func DisabledValidator() Validator           { var v *ruleValidator; return v }
