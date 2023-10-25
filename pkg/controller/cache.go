package controller

import (
	"sync"
)

// Cache is a simple in-memory cache.
type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

var cacheInstance *Cache
var once sync.Once

// Init initializes the cache.
func Init() {
	once.Do(func() {
		cacheInstance = &Cache{
			store: make(map[string]interface{}),
		}
	})
}

// GetCache returns the cache instance.
func GetCache() *Cache {
	return cacheInstance
}

// Set adds a key-value pair to the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

// Get retrieves the value associated with a key from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.store[key]
	return value, ok
}

// Delete removes a key-value pair from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}
