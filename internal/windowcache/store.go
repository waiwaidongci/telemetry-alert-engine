package windowcache

type Store struct{ values map[string]int }

func NewStore() *Store                    { return &Store{values: map[string]int{}} }
func (s *Store) Put(k string, v int)      { s.values[k] = v }
func (s *Store) Snapshot() map[string]int { return s.values }
