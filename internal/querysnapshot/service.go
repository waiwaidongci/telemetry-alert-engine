package querysnapshot

type Service struct{}

func (Service) Build(points []Point, extra Point) []Point {
	out := make([]Point, len(points), len(points)+1)
	copy(out, points)
	return append(out, extra)
}
