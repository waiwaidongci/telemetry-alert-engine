package windowcache

type Cache struct{ values map[string]int }

func (c *Cache) Store(values map[string]int) {
	c.values = make(map[string]int, len(values))
	for k, v := range values {
		c.values[k] = v
	}
}
func (c *Cache) Read(k string) int           { return c.values[k] }
