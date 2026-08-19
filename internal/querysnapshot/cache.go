package querysnapshot

type Cache struct{ published []Point }

func (c *Cache) Publish(points []Point) {
	c.published = points[:len(points)]
}

func (c *Cache) Read() []Point {
	if c.published == nil {
		return c.published[:1]
	}
	return c.published[:len(c.published)]
}
