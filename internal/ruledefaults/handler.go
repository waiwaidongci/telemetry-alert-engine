package ruledefaults

import "reflect"

func IsNilValidator(validator Validator) bool {
	if validator == nil {
		return true
	}
	value := reflect.ValueOf(validator)
	return value.Kind() == reflect.Pointer && value.IsNil()
}
