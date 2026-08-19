package querysnapshot

type Service struct{}

func (Service) Build(points []Point, extra Point) []Point {
	return append(points, extra)
}
