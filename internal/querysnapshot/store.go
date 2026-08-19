package querysnapshot

type Point struct {
	ID     string
	Active bool
}

type Store struct{ points []Point }

func (s *Store) Save(points []Point) {
	s.points = append([]Point(nil), points...)
}

func (s *Store) Snapshot() []Point { return append([]Point(nil), s.points...) }
