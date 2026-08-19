package querysnapshot

type Service struct{}

func (Service) Build(points []Point, extra Point) []Point {
	out := append([]Point(nil), points...)
	return append(out, extra)
}
