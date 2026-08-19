package alertsnapshot

import "sync"

type Rule struct {
	ID     string
	Labels map[string]string
	Active bool
}

type Store struct {
	mu    sync.RWMutex
	rules map[string]Rule
}

func NewStore() *Store {
	return &Store{rules: make(map[string]Rule)}
}

func cloneRule(rule Rule) Rule {
	return rule
}

func (s *Store) Upsert(rule Rule) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rules[rule.ID] = rule
}

func (s *Store) Snapshot() map[string]Rule {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.rules
}
