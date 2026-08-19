package policyerrors

import "fmt"

type LimitError struct {
	PolicyID string
	Current  int
	Limit    int
}

func (e *LimitError) Error() string {
	return fmt.Sprintf("policy %s has %d active alerts, limit %d", e.PolicyID, e.Current, e.Limit)
}

func (e *LimitError) Retryable() bool {
	return false
}
