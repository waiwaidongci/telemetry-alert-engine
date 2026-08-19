package ingestcontext

import "context"

type Worker struct{}

func (Worker) Retry(ctx context.Context, attempts int, call func(context.Context) error) error {
	var last error
	for attempt := 0; attempt < attempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		last = call(ctx)
		if last == nil {
			return nil
		}
	}
	return last
}
