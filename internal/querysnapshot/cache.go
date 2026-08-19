package querysnapshot

type Cache struct{ published []Point }

func (c *Cache) Publish(points []Point) {
	if len(points) == 0 {
		c.published = nil
		return
	}
	c.published = make([]Point, len(points))
	copy(c.published, points)
}

func (c *Cache) Read() []Point {
	if c.published == nil {
		return nil
	}
	return c.published[:len(c.published)]
}
