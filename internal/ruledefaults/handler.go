package ruledefaults

func IsNilValidator(validator Validator) bool {
	if validator == nil {
		return true
	}
	return false
}
