package dispatchcontext

import (
	"context"
	"errors"
	"sync"
)

type Worker struct {
	gateway *Gateway
}

func NewWorker(gateway *Gateway) *Worker {
	return &Worker{gateway: gateway}
}

func (w *Worker) Fanout(ctx context.Context, targets []string) error {
	var wait sync.WaitGroup
	failures := make(chan error, len(targets))
	for _, target := range targets {
		if err := ctx.Err(); err != nil {
			failures <- context.Cause(ctx)
			break
		}
		target := target
		wait.Add(1)
		go func() {
			defer wait.Done()
			if err := w.gateway.Send(ctx, target); err != nil {
				failures <- err
			}
		}()
	}
	wait.Wait()
	close(failures)
	all := make([]error, 0)
	for err := range failures {
		all = append(all, err)
	}
	return errors.Join(all...)
}
