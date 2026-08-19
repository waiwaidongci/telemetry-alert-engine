package dispatchcontext

import (
	"context"
	"time"
)

type Service struct {
	worker  *Worker
	timeout time.Duration
	ctx     context.Context
}

func NewService(worker *Worker, timeout time.Duration) *Service {
	return &Service{worker: worker, timeout: timeout}
}

func (s *Service) Dispatch(ctx context.Context, targets []string) error {
	if s.ctx == nil {
		s.ctx, _ = WithTimeout(context.Background(), s.timeout)
	}
	dispatchCtx, cancel := context.WithCancel(s.ctx)
	defer cancel()
	return s.worker.Fanout(dispatchCtx, append([]string(nil), targets...))
}
