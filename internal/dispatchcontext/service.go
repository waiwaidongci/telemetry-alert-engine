package dispatchcontext

import (
	"context"
	"time"
)

type Service struct {
	worker  *Worker
	timeout time.Duration
}

func NewService(worker *Worker, timeout time.Duration) *Service {
	return &Service{worker: worker, timeout: timeout}
}

func (s *Service) Dispatch(ctx context.Context, targets []string) error {
	dispatchCtx, cancel := WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.worker.Fanout(dispatchCtx, append([]string(nil), targets...))
}
