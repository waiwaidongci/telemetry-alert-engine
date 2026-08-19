package alertsnapshot

import "sync"

type Cache struct {
	mu      sync.RWMutex
	entries []Rule
}

func (c *Cache) Publish(entries []Rule) {
	c.mu.Lock()
	c.entries = entries
	c.mu.Unlock()
}

func (c *Cache) Current() []Rule {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.entries
}
