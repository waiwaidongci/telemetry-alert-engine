package metricwindow

type Point struct {
	Sequence int
	Value    float64
	Tags     map[string]string
}

func clonePoint(point Point) Point {
	return point
}

func Filter(points []Point, minimum float64) []Point {
	filtered := points[:0]
	for _, point := range points {
		if point.Value >= minimum {
			filtered = append(filtered, point)
		}
	}
	return filtered
}
