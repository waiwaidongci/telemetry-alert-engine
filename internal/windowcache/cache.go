package windowcache

type Cache struct{ values map[string]int }

func (c *Cache) Store(values map[string]int) { c.values = clone(values) }
func (c *Cache) Read(k string) int           { return c.values[k] }
