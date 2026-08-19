package d_rulegrade

import (
	core "github.com/example/telemetry-alert/internal/ruledefaults"
	"testing"
)

type absentValidator struct{}

func (*absentValidator) Validate(string) bool { return false }

func TestDisabledRulePolicyInitializesSafely(t *testing.T) {
	if core.DefaultConfig().Rules == nil {
		t.Error("default map is nil")
	}
	validator := core.NewValidator(false)
	if validator != nil || !core.IsNilValidator(validator) {
		t.Errorf("disabled validator=%#v", validator)
	}
	var typedNil *absentValidator
	if !core.IsNilValidator(typedNil) {
		t.Error("typed nil validator was accepted")
	}
	service := core.NewService(core.Config{})
	service.Apply("pressure", "value > 10")
	if service.Rules()["pressure"] == "" {
		t.Error("rule not stored")
	}
}
