package routingpolicy

import "testing"

func TestApplyDefaultRoutingPolicy(t *testing.T) {
	handler := NewHandler(NewService(DefaultConfig()), NewValidator(true))
	config, err := handler.Put("primary", "https://alerts.invalid")
	if err != nil {
		t.Fatalf("apply default policy: %v", err)
	}
	if config.Routes["primary"] == "" || config.Labels["last_route"] != "primary" {
		t.Fatalf("default maps were not populated: %#v", config)
	}
}

func TestDisabledValidatorIsSkipped(t *testing.T) {
	handler := NewHandler(NewService(Config{}), NewValidator(false))
	config, err := handler.Put("fallback", "https://fallback.invalid")
	if err != nil {
		t.Fatalf("disabled validator returned error: %v", err)
	}
	if config.Routes["fallback"] == "" {
		t.Fatal("route was not stored")
	}
}
