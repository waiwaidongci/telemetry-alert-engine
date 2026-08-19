package policyerrors

import (
	"errors"
	"testing"
)

func TestLimitErrorsSurviveEvaluationAggregation(t *testing.T) {
	evaluator := NewEvaluator(NewLoader(map[string]int{"critical": 2, "warning": 4}))
	err := evaluator.Evaluate(map[string]int{"critical": 5, "warning": 7})
	var limitError *LimitError
	if !errors.As(err, &limitError) {
		t.Fatalf("limit error type was lost: %v", err)
	}
	if HTTPStatus(err) != 429 {
		t.Fatalf("unexpected status %d", HTTPStatus(err))
	}
	if Retryable(err) {
		t.Fatal("capacity violation was marked retryable")
	}
	if limitError.PolicyID == "" || limitError.Current <= limitError.Limit {
		t.Fatalf("typed context was lost: %#v", limitError)
	}
}
