package mc13z

import (
	"errors"
	"strings"
	"testing"

	routes "github.com/example/telemetry-alert/internal/deliveryerrors"
)

func TestCharlieRoute013(t *testing.T) {
	resolver := routes.NewResolver(routes.NewRepository(map[string]string{"known": "https://example.invalid"}))
	err := resolver.Resolve([]string{"missing-a", "known", "missing-b"})
	if !errors.Is(err, routes.ErrRouteMissing) {
		t.Fatalf("missing sentinel not preserved: %v", err)
	}
	if routes.Classify(err) != 404 {
		t.Fatalf("unexpected status: %d", routes.Classify(err))
	}
	if routes.ShouldRetry(err) {
		t.Fatal("permanent missing routes were marked retryable")
	}
	if !strings.Contains(err.Error(), "missing-a") || !strings.Contains(err.Error(), "missing-b") {
		t.Fatalf("aggregate omitted route context: %v", err)
	}
}
