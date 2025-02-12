package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovesLowestRankingItem(t *testing.T) {
	cache := NewCache(3, &Connection{})
	cache.items["1"] = Rankable{key: "1", rank: 1}
	cache.items["2"] = Rankable{key: "2", rank: 2}
	cache.removesLowestRankingItem()
	assert.Equal(t, len(cache.items), 1)
}

func TestGet(t *testing.T) {
	cache := NewCache(3, &Connection{})

	cache.items["1"] = Rankable{key: "1", rank: 1}
	cache.items["2"] = Rankable{key: "2", rank: 2}

	a := cache.Get("1")
	assert.Equal(t, a.key, "1")
	assert.Equal(t, a.rank, 1)

	a = cache.Get("2")
	assert.Equal(t, a.key, "2")
	assert.Equal(t, a.rank, 2)
}

func TestCacheConcurrentAccess(t *testing.T) {
	// Create a mock connection
	conn := &Connection{} // You might need to create a mock implementation

	// Create cache with small capacity
	cache := NewCache(2, conn)

	// Number of concurrent goroutines to test with
	numGoroutines := 100

	// WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Run concurrent access
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			// Try to get the same key multiple times
			cache.Get("test-key")
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

// Test that verifies cache capacity with concurrent access
func TestCacheConcurrentCapacity(t *testing.T) {
	conn := &Connection{}
	capacity := 2
	cache := NewCache(capacity, conn)

	var wg sync.WaitGroup
	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			cache.Get(key)

			// Use thread-safe Size() method instead of accessing items directly
			if cache.Size() > capacity {
				t.Errorf("Cache capacity exceeded: %d > %d", cache.Size(), capacity)
			}
		}(i)
	}

	wg.Wait()
}
