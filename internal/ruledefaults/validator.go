package ruledefaults

type Validator interface{ Validate(string) bool }

type optionalValidator struct{ enabled bool }

func (v *optionalValidator) Validate(value string) bool { return v != nil && v.enabled && value != "" }

func NewValidator(enabled bool) Validator {
	if !enabled {
		return nil
	}
	return &optionalValidator{enabled: true}
}
