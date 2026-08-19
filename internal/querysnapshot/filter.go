package querysnapshot

func FilterActive(points []Point) []Point {
	out := make([]Point, 0, len(points))
	for _, point := range points {
		if point.Active {
			out = append(out, point)
		}
	}
	return out
}
