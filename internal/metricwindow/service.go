package metricwindow

type Service struct {
	store *Store
	cache *Cache
}

func NewService(store *Store, cache *Cache) *Service {
	return &Service{store: store, cache: cache}
}

func (s *Service) Build(deviceID string, source []Point, minimum float64) []Point {
	filtered := Filter(source, minimum)
	s.store.Save(deviceID, filtered)
	s.cache.Publish(filtered)
	return s.cache.Current()
}

func (s *Service) Extend(deviceID string, point Point) []Point {
	current := s.store.Load(deviceID)
	extended := make([]Point, 0, len(current)+1)
	extended = append(extended, current...)
	extended = append(extended, clonePoint(point))
	s.store.Save(deviceID, extended)
	s.cache.Publish(extended)
	return s.cache.Current()
}
