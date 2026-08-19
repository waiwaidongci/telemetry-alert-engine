package querysnapshot

func FilterActive(points []Point) []Point {
	out := points[:0]
	for _, point := range points {
		if point.Active {
			out = append(out, point)
		}
	}
	return out
}
