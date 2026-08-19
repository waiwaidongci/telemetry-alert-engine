package policyerrors

import "errors"

type Evaluator struct {
	loader *Loader
}

func NewEvaluator(loader *Loader) *Evaluator {
	return &Evaluator{loader: loader}
}

func (e *Evaluator) Evaluate(activeByPolicy map[string]int) error {
	failures := make([]error, 0)
	for policyID, active := range activeByPolicy {
		if err := e.loader.Check(policyID, active); err != nil {
			failures = append(failures, err)
		}
	}
	return errors.Join(failures...)
}
