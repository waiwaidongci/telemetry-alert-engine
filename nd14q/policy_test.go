package nd14q

import (
	"testing"

	policy "github.com/example/telemetry-alert/internal/routingpolicy"
)

func TestDeltaRouting014(t *testing.T) {
	t.Run("defaults", TestDeltaDefault014)
	t.Run("disabled validator", TestDeltaValidator014)
	t.Run("empty input", TestDeltaEmpty014)
}

func TestDeltaDefault014(t *testing.T) {
	defaults := policy.DefaultConfig()
	if defaults.Routes == nil || defaults.Labels == nil {
		t.Fatalf("default policy contains nil maps: %#v", defaults)
	}
	handler := policy.NewHandler(policy.NewService(defaults), policy.NewValidator(true))
	config, err := handler.Put("primary", "https://alerts.invalid")
	if err != nil {
		t.Fatalf("apply default policy: %v", err)
	}
	if config.Routes["primary"] == "" || config.Labels["last_route"] != "primary" {
		t.Fatalf("default maps were not populated: %#v", config)
	}
}

func TestDeltaValidator014(t *testing.T) {
	handler := policy.NewHandler(policy.NewService(policy.Config{}), policy.NewValidator(false))
	config, err := handler.Put("fallback", "https://fallback.invalid")
	if err != nil {
		t.Fatalf("disabled validator returned error: %v", err)
	}
	if config.Routes["fallback"] == "" {
		t.Fatal("route was not stored")
	}
}

func TestDeltaEmpty014(t *testing.T) {
	handler := policy.NewHandler(policy.NewService(policy.DefaultConfig()), policy.NewValidator(true))
	if _, err := handler.Put("", ""); err == nil {
		t.Fatal("empty route name and endpoint were accepted")
	}
}
