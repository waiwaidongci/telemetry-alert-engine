package querysnapshot

type Point struct {
	ID     string
	Active bool
}

type Store struct{ points []Point }

func (s *Store) Save(points []Point) {
	if len(points) == 0 {
		s.points = nil
		return
	}
	s.points = make([]Point, len(points))
	copy(s.points, points)
}

func (s *Store) Snapshot() []Point {
	if s.points == nil {
		return nil
	}
	return s.points[:len(s.points)]
}
