package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) *Cache {
	c := Cache{
		data:     map[string]cacheEntry{},
		interval: t,
	}
	ticker := time.NewTicker(t)

	go func(*Cache) {
		for {
			<-ticker.C
			c.ReapLoop()
		}
	}(&c)

	return &c
}

func (c *Cache) Add(key string, entry []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       entry,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.data[key]; ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) ReapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.data {
		if time.Since(entry.createdAt) >= c.interval {
			delete(c.data, key)
		}
	}
}
