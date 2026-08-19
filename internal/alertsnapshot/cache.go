package alertsnapshot

import "sync"

type Cache struct {
	mu      sync.RWMutex
	entries []Rule
}

func (c *Cache) Publish(entries []Rule) {
	cloned := make([]Rule, len(entries))
	for index, entry := range entries {
		cloned[index] = cloneRule(entry)
	}
	c.mu.Lock()
	c.entries = cloned
	c.mu.Unlock()
}

func (c *Cache) Current() []Rule {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cloned := make([]Rule, len(c.entries))
	for index, entry := range c.entries {
		cloned[index] = cloneRule(entry)
	}
	return cloned
}
