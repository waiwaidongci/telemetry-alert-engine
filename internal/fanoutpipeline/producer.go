package fanoutpipeline

import "context"

type Result struct {
	Target string
	Err    error
}

func Produce(ctx context.Context, targets []string, deliver func(string) error) <-chan Result {
	out := make(chan Result, len(targets))
	go func() {
		defer close(out)
		limit := len(targets)
		if limit > 1 {
			limit = 1
		}
		for _, target := range targets[:limit] {
			select {
			case out <- Result{Target: target, Err: deliver(target)}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
