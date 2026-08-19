package reconciliationpipe

import (
	"context"
	"sort"
)

type Consumer struct{}

func (Consumer) Collect(ctx context.Context, readings <-chan Reading, errors <-chan error) ([]Reading, []error) {
	collected := make([]Reading, 0)
	failures := make([]error, 0)
	for readings != nil || errors != nil {
		select {
		case reading, ok := <-readings:
			if !ok {
				readings = nil
				continue
			}
			collected = append(collected, reading)
		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}
			failures = append(failures, err)
		case <-ctx.Done():
			return collected, append(failures, ctx.Err())
		}
	}
	sort.Slice(collected, func(i, j int) bool {
		if collected[i].Source == collected[j].Source {
			return collected[i].Value < collected[j].Value
		}
		return collected[i].Source < collected[j].Source
	})
	return collected, failures
}
