package sweepcontext

import "context"

type Service struct {
	repository *Repository
	worker     *Worker
}

func NewService(repository *Repository, worker *Worker) *Service {
	return &Service{repository: repository, worker: worker}
}

func (s *Service) Sweep(ctx context.Context) error {
	return s.repository.Load(ctx)
}

func (s *Service) RetrySweep(ctx context.Context) error {
	return s.worker.Retry(ctx, s.repository.Load)
}
