package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
	interval	time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(t time.Duration) *Cache {
	c := Cache{
		data: map[string]cacheEntry{},
		interval: t,
	}
	return &c
}

func (c *Cache) Add(key string, entry []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: entry,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte,bool) {
	c.mu.Lock()
	if entry,ok := c.data[key]; ok {
		return entry.val,true
	}
	c.mu.Unlock()
	return nil,false
}

func (c *Cache) ReapLoop() {
	c.mu.Lock()
	for key,entry := range c.data {
		if time.Since(entry.createdAt) > c.interval { 
		delete(c.data,key)
		}
	}
	c.mu.Unlock()
}