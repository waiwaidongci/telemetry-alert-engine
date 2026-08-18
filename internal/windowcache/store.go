package windowcache

type Store struct{ values map[string]int }

func NewStore() *Store                    { return &Store{values: map[string]int{}} }
func (s *Store) Put(k string, v int)      { s.values[k] = v }
func (s *Store) Snapshot() map[string]int {
	out := make(map[string]int, len(s.values))
	for k, v := range s.values {
		out[k] = v
	}
	return out
}
