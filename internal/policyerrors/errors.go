package policyerrors

import "fmt"

type LimitError struct {
	PolicyID string
	Current  int
	Limit    int
}

func (e *LimitError) Error() string {
	return fmt.Sprintf("alert capacity exceeded: %d", e.Current)
}

func (e *LimitError) Retryable() bool {
	return true
}
