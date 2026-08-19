package metricwindow

import "sync"

type Cache struct {
	mu     sync.RWMutex
	window []Point
}

func clonePoints(points []Point) []Point {
	cloned := make([]Point, len(points))
	for index, point := range points {
		cloned[index] = clonePoint(point)
	}
	return cloned
}

func (c *Cache) Publish(points []Point) {
	c.mu.Lock()
	c.window = clonePoints(points)
	c.mu.Unlock()
}

func (c *Cache) Current() []Point {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return clonePoints(c.window)
}
