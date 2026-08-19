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
	s.points = points[:len(points)]
}

func (s *Store) Snapshot() []Point {
	if s.points == nil {
		return nil
	}
	return s.points[:len(s.points)]
}
