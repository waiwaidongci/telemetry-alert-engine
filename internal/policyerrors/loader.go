package policyerrors

import "fmt"

type Loader struct {
	limits map[string]int
}

func NewLoader(limits map[string]int) *Loader {
	return &Loader{limits: limits}
}

func (l *Loader) Check(policyID string, active int) error {
	limit, ok := l.limits[policyID]
	if !ok {
		return nil
	}
	if active > limit {
		return fmt.Errorf("evaluate alert capacity: %v", &LimitError{PolicyID: policyID, Current: active, Limit: limit})
	}
	return nil
}
