package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that verifies removeLowestRankingItem handles empty cache gracefully
func TestRemoveLowestRankingItemEmptyCache(t *testing.T) {
	cache := NewCache(3, &Connection{})

	// Verify cache is empty
	assert.Equal(t, 0, cache.Size())

	// This should NOT panic - it should handle empty cache gracefully
	cache.removeLowestRankingItem()

	// Cache should still be empty
	assert.Equal(t, 0, cache.Size())
}

// Test that verifies correct item is removed (lowest rank)
func TestRemoveLowestRankingItemCorrectEviction(t *testing.T) {
	cache := NewCache(3, &Connection{})

	// Add items with known ranks
	cache.items["high"] = Rankable{key: "high", rank: 100}
	cache.items["medium"] = Rankable{key: "medium", rank: 50}
	cache.items["low"] = Rankable{key: "low", rank: 10}

	assert.Equal(t, 3, cache.Size())

	// Remove lowest ranking item
	cache.removeLowestRankingItem()

	// Should have 2 items left
	assert.Equal(t, 2, cache.Size())

	// "low" (rank 10) should be gone
	_, exists := cache.items["low"]
	assert.False(t, exists)

	// "medium" and "high" should still exist
	_, exists = cache.items["medium"]
	assert.True(t, exists)
	_, exists = cache.items["high"]
	assert.True(t, exists)
}

// Test that verifies single item removal
func TestRemoveLowestRankingItemSingleItem(t *testing.T) {
	cache := NewCache(3, &Connection{})

	cache.items["only"] = Rankable{key: "only", rank: 42}
	assert.Equal(t, 1, cache.Size())

	cache.removeLowestRankingItem()

	assert.Equal(t, 0, cache.Size())
	_, exists := cache.items["only"]
	assert.False(t, exists)
}

// Test tie-breaking behavior when multiple items have same rank
func TestRemoveLowestRankingItemTieBreaking(t *testing.T) {
	cache := NewCache(3, &Connection{})

	// Add multiple items with same lowest rank
	cache.items["tie1"] = Rankable{key: "tie1", rank: 10}
	cache.items["tie2"] = Rankable{key: "tie2", rank: 10}
	cache.items["high"] = Rankable{key: "high", rank: 50}

	assert.Equal(t, 3, cache.Size())

	cache.removeLowestRankingItem()

	// Should remove one of the rank-10 items (non-deterministic which one)
	assert.Equal(t, 2, cache.Size())

	// At least one rank-10 item should be gone
	tie1Exists := false
	tie2Exists := false
	if _, ok := cache.items["tie1"]; ok {
		tie1Exists = true
	}
	if _, ok := cache.items["tie2"]; ok {
		tie2Exists = true
	}

	// Exactly one of the ties should remain
	assert.True(t, tie1Exists != tie2Exists, "Exactly one tie item should remain")

	// High should always remain
	_, exists := cache.items["high"]
	assert.True(t, exists)
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
