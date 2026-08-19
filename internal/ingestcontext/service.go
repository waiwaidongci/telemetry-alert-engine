package ingestcontext

import "context"

type Service struct {
	gateway Gateway
	worker  Worker
}

func NewService() Service { return Service{gateway: Gateway{}, worker: Worker{}} }

func (s Service) Ingest(ctx context.Context, ready <-chan struct{}, call func(context.Context) error) error {
	request := NewRequest(ctx)
	if err := s.gateway.Load(request.Context(), ready); err != nil {
		return err
	}
	return s.worker.Retry(request.Context(), 3, call)
}
