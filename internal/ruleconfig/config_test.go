package ruleconfig

import "testing"

func TestDefaultRuleConfigHandlesZeroValues(t *testing.T) {
	c := LoadDefaults()
	if c.Rules == nil {
		t.Error("default rules map is nil")
	}
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("apply panicked: %v", r)
			}
		}()
		Apply(&c, "temperature", 80)
	}()
	if DisabledValidator() != nil {
		t.Error("disabled validator is a typed nil")
	}
	var concrete *ruleValidator
	var v Validator = concrete
	if !IsNilValidator(v) {
		t.Error("typed nil validator was treated as usable")
	}
}
