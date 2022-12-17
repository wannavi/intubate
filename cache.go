package intubate

import "sync"

type Cache struct {
	context map[string]interface{}
	mutex   sync.Mutex
}

func NewCache() *Cache {
	cache := &Cache{
		context: make(map[string]interface{}),
		mutex:   sync.Mutex{},
	}

	return cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	c.context[key] = value
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if v, ok := c.context[key]; ok {
		return v
	}

	return nil
}
