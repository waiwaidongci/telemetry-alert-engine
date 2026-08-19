package fanoutpipeline

import (
	"context"
	"sync"
)

type Result struct {
	Target string
	Err    error
}

// Produce launches one goroutine per target, delivering concurrently. The
// returned channel is buffered to len(targets) so that successful sends never
// block, and it is closed once every goroutine finishes. When ctx is cancelled,
// goroutines that have not yet sent their result abandon the send and return.
func Produce(ctx context.Context, targets []string, deliver func(string) error) <-chan Result {
	out := make(chan Result, len(targets))
	var wg sync.WaitGroup
	for _, target := range targets {
		wg.Add(1)
		target := target
		go func() {
			defer wg.Done()
			err := deliver(target)
			select {
			case out <- Result{Target: target, Err: err}:
			case <-ctx.Done():
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
