package fanoutpipeline

import "context"

// Consume reads results until the channel is closed or ctx is cancelled. On
// cancellation it returns the context error so callers can distinguish a
// complete fan-out from an aborted one.
func Consume(ctx context.Context, results <-chan Result) ([]Result, error) {
	var collected []Result
	for {
		select {
		case result, ok := <-results:
			if !ok {
				return collected, nil
			}
			collected = append(collected, result)
		case <-ctx.Done():
			return collected, ctx.Err()
		}
	}
}
