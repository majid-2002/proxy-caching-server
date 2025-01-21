package cache

import "sync"

type Cache struct {
	storage map[string]interface{}
	mu      sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage = make(map[string]interface{})
}
