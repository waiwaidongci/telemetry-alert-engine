package ruledefaults

import "testing"

func TestRuleZeroValueInitializesPolicyAndRejectsTypedNil(t *testing.T) {
	config := DefaultConfig()
	if config.Rules == nil {
		t.Error("default config left nil rules")
	}
	if validator := NewValidator(false); validator != nil {
		t.Errorf("disabled factory returned typed nil: %#v", validator)
	}
	service := NewService(Config{})
	service.Apply("pressure", "value > 10")
	if service.Rules()["pressure"] == "" {
		t.Error("service did not initialize and store rule")
	}
	var typedNil *optionalValidator
	if !IsNilValidator(typedNil) {
		t.Error("handler accepted typed nil validator")
	}
}
