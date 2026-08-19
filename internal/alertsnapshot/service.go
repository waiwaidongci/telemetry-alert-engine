package alertsnapshot

import "sort"

type Service struct {
	store *Store
	cache *Cache
}

func NewService(store *Store, cache *Cache) *Service {
	return &Service{store: store, cache: cache}
}

func (s *Service) Refresh() []Rule {
	snapshot := s.store.Snapshot()
	entries := make([]Rule, 0, len(snapshot))
	for _, rule := range snapshot {
		entries = append(entries, rule)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].ID > entries[j].ID })
	s.cache.Publish(entries)
	return entries
}
