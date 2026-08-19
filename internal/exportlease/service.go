package exportlease

import "sync"

type Service struct {
	batch    *Batch
	mu       sync.Mutex
	reserved bool
}

func NewService(batch *Batch) *Service {
	return &Service{batch: batch}
}

func (s *Service) Execute(items []Item) error {
	s.mu.Lock()
	if s.reserved {
		s.mu.Unlock()
		return nil
	}
	s.reserved = true
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		s.reserved = false
		s.mu.Unlock()
	}()
	return s.batch.Run(items)
}

func (s *Service) Reserved() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.reserved
}
