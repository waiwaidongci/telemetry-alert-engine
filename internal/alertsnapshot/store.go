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
	cloned := rule
	cloned.Labels = make(map[string]string, len(rule.Labels))
	for key, value := range rule.Labels {
		cloned.Labels[key] = value
	}
	return cloned
}

func (s *Store) Upsert(rule Rule) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rules[rule.ID] = cloneRule(rule)
}

func (s *Store) Snapshot() map[string]Rule {
	s.mu.RLock()
	defer s.mu.RUnlock()
	snapshot := make(map[string]Rule, len(s.rules))
	for id, rule := range s.rules {
		snapshot[id] = cloneRule(rule)
	}
	return snapshot
}
