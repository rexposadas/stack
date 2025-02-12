package main

import "sync"

type Rankable struct {
	key  string
	rank int
}

type Cache struct {
	capacity int

	// key is the key of the the Rankable item
	items      map[string]Rankable
	connection *Connection
	mutex      sync.Mutex
}

func NewCache(capacity int, connection *Connection) *Cache {
	return &Cache{
		capacity:   capacity,
		connection: connection,
		items:      make(map[string]Rankable),
	}
}

func (c *Cache) removesLowestRankingItem() {
	// find the lowest ranking item
	var lowestRankingItem Rankable
	iter := 0
	for _, item := range c.items {
		if iter == 0 {
			lowestRankingItem = item
			iter++
			continue
		}
		if item.rank < lowestRankingItem.rank {
			lowestRankingItem = item
		}
	}

	delete(c.items, lowestRankingItem.key)
}

func (c *Cache) Get(key string) Rankable {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// check if item is in the cache (items)
	if item, ok := c.items[key]; ok {
		return item
	}

	item := c.connection.getItem(key)

	// add item to the cache
	if len(c.items) < c.capacity {
		c.items[item.key] = item
		return item
	}

	// At this point, we are at cache capacity.

	// step 1: remove lowest rank item
	c.removesLowestRankingItem()

	// step 2: add the item search for.
	c.items[item.key] = item

	return item
}

// Add this new method to safely get the cache size
func (c *Cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return len(c.items)
}
