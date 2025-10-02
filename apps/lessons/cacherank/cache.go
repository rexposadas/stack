package main

import (
	"errors"
	"sync"
)

var (
	ErrInvalidCapacity = errors.New("capacity must be greater than 0")
	ErrConnectionNil   = errors.New("connection cannot be nil")
)

type Rankable struct {
	Key  string
	Rank int
}

// Stats holds cache performance metrics
type Stats struct {
	Hits      int64 // Number of cache hits
	Misses    int64 // Number of cache misses
	Evictions int64 // Number of evictions
}

type Cache struct {
	capacity int

	// items maps keys to Rankable items
	items      map[string]Rankable
	connection *Connection
	mutex      sync.Mutex

	// metrics
	stats Stats
}

func NewCache(capacity int, connection *Connection) (*Cache, error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}
	if connection == nil {
		return nil, ErrConnectionNil
	}
	return &Cache{
		capacity:   capacity,
		connection: connection,
		items:      make(map[string]Rankable),
	}, nil
}

func (c *Cache) removeLowestRankingItem() {
	if len(c.items) == 0 {
		return
	}

	var lowestRankingItem Rankable
	var found bool

	for _, item := range c.items {
		if !found || item.Rank < lowestRankingItem.Rank {
			lowestRankingItem = item
			found = true
		}
	}

	delete(c.items, lowestRankingItem.Key)
	c.stats.Evictions++
}

func (c *Cache) Get(key string) (Rankable, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if item is in the cache
	if item, ok := c.items[key]; ok {
		c.stats.Hits++
		return item, nil
	}

	// Cache miss
	c.stats.Misses++

	// Fetch from connection
	item, err := c.connection.getItem(key)
	if err != nil {
		return Rankable{}, err
	}

	// Add item to the cache
	if len(c.items) < c.capacity {
		c.items[item.Key] = item
		return item, nil
	}

	// At this point, we are at cache capacity

	// Step 1: remove lowest rank item
	c.removeLowestRankingItem()

	// Step 2: add the item searched for
	c.items[item.Key] = item

	return item, nil
}

// Size returns the current number of items in the cache
func (c *Cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return len(c.items)
}

// Put adds or updates an item in the cache with explicit rank
func (c *Cache) Put(key string, rank int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item := Rankable{Key: key, Rank: rank}

	// If cache is not full, just add
	if len(c.items) < c.capacity {
		c.items[key] = item
		return
	}

	// If key already exists, update in place
	if _, exists := c.items[key]; exists {
		c.items[key] = item
		return
	}

	// Cache is full and key doesn't exist - evict lowest rank
	c.removeLowestRankingItem()
	c.items[key] = item
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.items[key]; exists {
		delete(c.items, key)
		return true
	}
	return false
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items = make(map[string]Rankable)
}

// Contains checks if a key exists in the cache
func (c *Cache) Contains(key string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exists := c.items[key]
	return exists
}

// Keys returns all keys currently in the cache
func (c *Cache) Keys() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	keys := make([]string, 0, len(c.items))
	for k := range c.items {
		keys = append(keys, k)
	}
	return keys
}

// GetStats returns a copy of the current cache statistics
func (c *Cache) GetStats() Stats {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.stats
}

// ResetStats resets all cache statistics to zero
func (c *Cache) ResetStats() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.stats = Stats{}
}
