package cache

import (
	"fmt"
	"sync"
)

func newMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[string][]byte),
		mutex: sync.RWMutex{},
	}
}

type MemoryCache struct {
	cache map[string][]byte // the built-in map in the SDK is used as the carrier of the memory cache, but it is not safe
	mutex sync.RWMutex      // select sync.RWMutex to ensure that the map is concurrently safe
	Stat
}

func (c *MemoryCache) Set(key string, value []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	old, existed := c.cache[key]
	if existed {
		c.del(key, old)
	}
	c.cache[key] = value
	c.add(key, value)
	return nil
}

func (c *MemoryCache) Get(key string) (value []byte, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.cache[key], nil
}

func (c *MemoryCache) Del(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, existed := c.cache[key]
	if !existed {
		return fmt.Errorf("%s not found", key)
	}
	delete(c.cache, key)
	c.del(key, value)
	return nil
}

func (c *MemoryCache) GetStat() Stat {
	return c.Stat
}
