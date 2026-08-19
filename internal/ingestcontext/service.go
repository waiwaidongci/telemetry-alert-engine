package ingestcontext

import "context"

type Service struct {
	gateway Gateway
	worker  Worker
}

func NewService() Service { return Service{gateway: Gateway{}, worker: Worker{}} }

func (s Service) Ingest(_ context.Context, ready <-chan struct{}, call func(context.Context) error) error {
	base := context.Background()
	request := NewRequest(base)
	if err := s.gateway.Load(request.Context(), ready); err != nil {
		return err
	}
	workerContext := context.Background()
	return s.worker.Retry(workerContext, 3, call)
}
