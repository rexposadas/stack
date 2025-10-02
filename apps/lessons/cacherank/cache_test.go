package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that verifies removeLowestRankingItem handles empty cache gracefully
func TestRemoveLowestRankingItemEmptyCache(t *testing.T) {
	cache, err := NewCache(3, &Connection{})
	assert.NoError(t, err)

	// Verify cache is empty
	assert.Equal(t, 0, cache.Size())

	// This should NOT panic - it should handle empty cache gracefully
	cache.removeLowestRankingItem()

	// Cache should still be empty
	assert.Equal(t, 0, cache.Size())
}

// Test that verifies correct item is removed (lowest rank)
func TestRemoveLowestRankingItemCorrectEviction(t *testing.T) {
	cache, err := NewCache(3, &Connection{})
	assert.NoError(t, err)

	// Add items with known ranks using Put
	cache.Put("high", 100)
	cache.Put("medium", 50)
	cache.Put("low", 10)

	assert.Equal(t, 3, cache.Size())

	// Remove lowest ranking item
	cache.removeLowestRankingItem()

	// Should have 2 items left
	assert.Equal(t, 2, cache.Size())

	// "low" (rank 10) should be gone
	assert.False(t, cache.Contains("low"))

	// "medium" and "high" should still exist
	assert.True(t, cache.Contains("medium"))
	assert.True(t, cache.Contains("high"))
}

// Test that verifies single item removal
func TestRemoveLowestRankingItemSingleItem(t *testing.T) {
	cache, err := NewCache(3, &Connection{})
	assert.NoError(t, err)

	cache.Put("only", 42)
	assert.Equal(t, 1, cache.Size())

	cache.removeLowestRankingItem()

	assert.Equal(t, 0, cache.Size())
	assert.False(t, cache.Contains("only"))
}

// Test tie-breaking behavior when multiple items have same rank
func TestRemoveLowestRankingItemTieBreaking(t *testing.T) {
	cache, err := NewCache(3, &Connection{})
	assert.NoError(t, err)

	// Add multiple items with same lowest rank
	cache.Put("tie1", 10)
	cache.Put("tie2", 10)
	cache.Put("high", 50)

	assert.Equal(t, 3, cache.Size())

	cache.removeLowestRankingItem()

	// Should remove one of the rank-10 items (non-deterministic which one)
	assert.Equal(t, 2, cache.Size())

	// Exactly one of the ties should remain
	tie1Exists := cache.Contains("tie1")
	tie2Exists := cache.Contains("tie2")
	assert.True(t, tie1Exists != tie2Exists, "Exactly one tie item should remain")

	// High should always remain
	assert.True(t, cache.Contains("high"))
}

func TestGet(t *testing.T) {
	cache, err := NewCache(3, &Connection{})
	assert.NoError(t, err)

	cache.Put("1", 1)
	cache.Put("2", 2)

	a, err := cache.Get("1")
	assert.NoError(t, err)
	assert.Equal(t, "1", a.Key)
	assert.Equal(t, 1, a.Rank)

	a, err = cache.Get("2")
	assert.NoError(t, err)
	assert.Equal(t, "2", a.Key)
	assert.Equal(t, 2, a.Rank)
}

func TestCacheConcurrentAccess(t *testing.T) {
	conn := &Connection{}
	cache, err := NewCache(2, conn)
	assert.NoError(t, err)

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
			_, _ = cache.Get("test-key")
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

// Test that verifies cache capacity with concurrent access
func TestCacheConcurrentCapacity(t *testing.T) {
	conn := &Connection{}
	capacity := 2
	cache, err := NewCache(capacity, conn)
	assert.NoError(t, err)

	var wg sync.WaitGroup
	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			_, _ = cache.Get(key)

			// Use thread-safe Size() method instead of accessing items directly
			if cache.Size() > capacity {
				t.Errorf("Cache capacity exceeded: %d > %d", cache.Size(), capacity)
			}
		}(i)
	}

	wg.Wait()
}
