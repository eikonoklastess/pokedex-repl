package pkcache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mux          *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mux.Unlock()
}

func (c *Cache) Get(key *string) ([]byte, bool) {
	if key == nil {
		return nil, false
	}
	c.mux.Lock()
	entry, ok := c.cacheEntries[*key]
	c.mux.Unlock()
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.After(interval)
		for val, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > interval {
				c.mux.Lock()
				delete(c.cacheEntries, val)
				c.mux.Unlock()
			}
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mux:          &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}
