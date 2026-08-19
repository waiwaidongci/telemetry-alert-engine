package querysnapshot

type Cache struct{ published []Point }

func (c *Cache) Publish(points []Point) { c.published = append([]Point(nil), points...) }

func (c *Cache) Read() []Point { return append([]Point(nil), c.published...) }
