package deliveryerrors

import (
	"errors"
	"strings"
	"testing"
)

func TestMissingRoutesKeepIdentityAcrossAggregation(t *testing.T) {
	resolver := NewResolver(NewRepository(map[string]string{"known": "https://example.invalid"}))
	err := resolver.Resolve([]string{"missing-a", "known", "missing-b"})
	if !errors.Is(err, ErrRouteMissing) {
		t.Fatalf("missing sentinel not preserved: %v", err)
	}
	if Classify(err) != 404 {
		t.Fatalf("unexpected status: %d", Classify(err))
	}
	if ShouldRetry(err) {
		t.Fatal("permanent missing routes were marked retryable")
	}
	if !strings.Contains(err.Error(), "missing-a") || !strings.Contains(err.Error(), "missing-b") {
		t.Fatalf("aggregate omitted route context: %v", err)
	}
}
