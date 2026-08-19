package metricwindow

import "sync"

type Cache struct {
	mu     sync.RWMutex
	window []Point
}

func clonePoints(points []Point) []Point {
	return points
}

func (c *Cache) Publish(points []Point) {
	c.mu.Lock()
	c.window = points
	c.mu.Unlock()
}

func (c *Cache) Current() []Point {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.window
}
