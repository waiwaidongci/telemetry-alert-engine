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
		w.attempts.Add(1)
		if err := operation(context.Background()); err == nil {
			return nil
		}
		time.Sleep(w.backoff)
	}
}

func (w *Worker) Attempts() int64 {
	return w.attempts.Load()
}
