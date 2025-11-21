package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(d map[string]cacheEntry, m sync.Mutex) *Cache {
	new_cache := Cache{
		data: d,
		mu: m,
	}
	return &new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}