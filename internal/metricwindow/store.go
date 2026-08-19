package metricwindow

import "sync"

type Store struct {
	mu      sync.RWMutex
	windows map[string][]Point
}

func NewStore() *Store {
	return &Store{windows: make(map[string][]Point)}
}

func (s *Store) Save(deviceID string, points []Point) {
	s.mu.Lock()
	s.windows[deviceID] = points
	s.mu.Unlock()
}

func (s *Store) Load(deviceID string) []Point {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.windows[deviceID]
}
