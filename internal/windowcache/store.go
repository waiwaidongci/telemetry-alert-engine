package windowcache

import "sync"

type Store struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewStore() *Store                    { return &Store{values: map[string]int{}} }
func (s *Store) Put(k string, v int)      { s.mu.Lock(); defer s.mu.Unlock(); s.values[k] = v }
func (s *Store) Snapshot() map[string]int { s.mu.RLock(); defer s.mu.RUnlock(); return clone(s.values) }
func clone(in map[string]int) map[string]int {
	out := make(map[string]int, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
