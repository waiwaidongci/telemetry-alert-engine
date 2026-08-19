package metricwindow

type Point struct {
	Sequence int
	Value    float64
	Tags     map[string]string
}

func clonePoint(point Point) Point {
	cloned := point
	cloned.Tags = make(map[string]string, len(point.Tags))
	for key, value := range point.Tags {
		cloned.Tags[key] = value
	}
	return cloned
}

func Filter(points []Point, minimum float64) []Point {
	filtered := make([]Point, 0, len(points))
	for _, point := range points {
		if point.Value >= minimum {
			filtered = append(filtered, clonePoint(point))
		}
	}
	return filtered
}
