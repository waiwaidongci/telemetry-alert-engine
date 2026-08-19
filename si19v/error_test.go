package si19v

import (
	"errors"
	"testing"

	policy "github.com/example/telemetry-alert/internal/policyerrors"
)

func TestIndiaPolicy019(t *testing.T) {
	evaluator := policy.NewEvaluator(policy.NewLoader(map[string]int{"critical": 2, "warning": 4}))
	err := evaluator.Evaluate(map[string]int{"critical": 5, "warning": 7})
	var limitError *policy.LimitError
	if !errors.As(err, &limitError) {
		t.Fatalf("limit error type was lost: %v", err)
	}
	if policy.HTTPStatus(err) != 429 {
		t.Fatalf("unexpected status %d", policy.HTTPStatus(err))
	}
	if policy.Retryable(err) {
		t.Fatal("capacity violation was marked retryable")
	}
	if limitError.PolicyID == "" || limitError.Current <= limitError.Limit {
		t.Fatalf("typed context was lost: %#v", limitError)
	}
}
