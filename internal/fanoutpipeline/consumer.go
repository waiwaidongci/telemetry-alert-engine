package fanoutpipeline

import "context"

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
