package sweepcontext

import (
	"context"
	"sync/atomic"
	"time"
)

type Worker struct {
	attempts atomic.Int64
	backoff  time.Duration
}

func NewWorker(backoff time.Duration) *Worker {
	return &Worker{backoff: backoff}
}

func (w *Worker) Retry(ctx context.Context, operation func(context.Context) error) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		w.attempts.Add(1)
		if err := operation(ctx); err == nil {
			return nil
		}
		timer := time.NewTimer(w.backoff)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		}
	}
}

func (w *Worker) Attempts() int64 {
	return w.attempts.Load()
}
