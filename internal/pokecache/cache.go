package pokecache

import (
	"sync"
	"time"
)

// Cache stores keyed data in memory for faster access.
// A mutex protects entries from concurrent access.
type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

// cacheEntry holds a value and the time it was created.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache creates a cache and starts background cleanup.
// The duration sets both the cleanup interval and the age
// after which entries are eligible for removal.
func NewCache(duration time.Duration) *Cache {
	newC := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: duration,
	}
	go newC.reapLoop()

	return newC

}

// reapLoop periodically removes expired entries.
// The mutex prevents concurrent map access during cleanup.
func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		cache.mu.Lock()
		for key, value := range cache.entries {
			if time.Since(value.createdAt) >= cache.interval {
				delete(cache.entries, key)

			}
		}
		cache.mu.Unlock()
	}

}

// Add stores a value under key, replacing any existing entry
// and resetting its creation time.
// The mutex protects the map during the update.
func (cache *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.mu.Lock()
	cache.entries[key] = entry
	cache.mu.Unlock()
}

// Get returns the value associated with key and whether it was found.
// The mutex protects the map during the lookup.
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	entry, ok := cache.entries[key]
	cache.mu.Unlock()
	return entry.val, ok

}
