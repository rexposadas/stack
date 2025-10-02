package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCacheValidation(t *testing.T) {
	// Test invalid capacity
	_, err := NewCache(0, &Connection{})
	assert.Equal(t, ErrInvalidCapacity, err)

	_, err = NewCache(-5, &Connection{})
	assert.Equal(t, ErrInvalidCapacity, err)

	// Test nil connection
	_, err = NewCache(10, nil)
	assert.Equal(t, ErrConnectionNil, err)

	// Test valid cache
	cache, err := NewCache(10, &Connection{})
	assert.NoError(t, err)
	assert.NotNil(t, cache)
}

func TestPut(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	// Add items
	cache.Put("key1", 10)
	cache.Put("key2", 20)
	cache.Put("key3", 30)

	assert.Equal(t, 3, cache.Size())
	assert.True(t, cache.Contains("key1"))
	assert.True(t, cache.Contains("key2"))
	assert.True(t, cache.Contains("key3"))
}

func TestPutUpdate(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	cache.Put("key1", 10)
	assert.True(t, cache.Contains("key1"))

	// Update existing key
	cache.Put("key1", 99)
	assert.Equal(t, 1, cache.Size())
	assert.True(t, cache.Contains("key1"))
}

func TestPutEviction(t *testing.T) {
	cache, _ := NewCache(2, &Connection{})

	cache.Put("low", 10)
	cache.Put("high", 100)
	assert.Equal(t, 2, cache.Size())

	// This should evict "low" (rank 10)
	cache.Put("medium", 50)
	assert.Equal(t, 2, cache.Size())
	assert.False(t, cache.Contains("low"))
	assert.True(t, cache.Contains("high"))
	assert.True(t, cache.Contains("medium"))
}

func TestDelete(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	cache.Put("key1", 10)
	assert.True(t, cache.Contains("key1"))

	// Delete existing key
	deleted := cache.Delete("key1")
	assert.True(t, deleted)
	assert.False(t, cache.Contains("key1"))
	assert.Equal(t, 0, cache.Size())

	// Delete non-existent key
	deleted = cache.Delete("nonexistent")
	assert.False(t, deleted)
}

func TestClear(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	cache.Put("key1", 10)
	cache.Put("key2", 20)
	cache.Put("key3", 30)
	assert.Equal(t, 3, cache.Size())

	cache.Clear()
	assert.Equal(t, 0, cache.Size())
	assert.False(t, cache.Contains("key1"))
	assert.False(t, cache.Contains("key2"))
	assert.False(t, cache.Contains("key3"))
}

func TestContains(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	assert.False(t, cache.Contains("key1"))

	cache.Put("key1", 10)
	assert.True(t, cache.Contains("key1"))

	cache.Delete("key1")
	assert.False(t, cache.Contains("key1"))
}

func TestKeys(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	// Empty cache
	keys := cache.Keys()
	assert.Equal(t, 0, len(keys))

	// Add some items
	cache.Put("key1", 10)
	cache.Put("key2", 20)
	cache.Put("key3", 30)

	keys = cache.Keys()
	assert.Equal(t, 3, len(keys))
	assert.Contains(t, keys, "key1")
	assert.Contains(t, keys, "key2")
	assert.Contains(t, keys, "key3")
}

func TestGetStats(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	// Initial stats
	stats := cache.GetStats()
	assert.Equal(t, int64(0), stats.Hits)
	assert.Equal(t, int64(0), stats.Misses)
	assert.Equal(t, int64(0), stats.Evictions)

	// Add an item and access it
	cache.Put("key1", 10)
	_, _ = cache.Get("key1") // Hit
	_, _ = cache.Get("key2") // Miss

	stats = cache.GetStats()
	assert.Equal(t, int64(1), stats.Hits)
	assert.Equal(t, int64(1), stats.Misses)

	// Trigger eviction
	cache.Put("key2", 20)
	cache.Put("key3", 30)
	cache.Put("key4", 40) // This should evict key1 (lowest rank)

	stats = cache.GetStats()
	assert.Equal(t, int64(1), stats.Evictions)
}

func TestResetStats(t *testing.T) {
	cache, _ := NewCache(3, &Connection{})

	cache.Put("key1", 10)
	_, _ = cache.Get("key1")
	_, _ = cache.Get("key2")

	stats := cache.GetStats()
	assert.Greater(t, stats.Hits, int64(0))
	assert.Greater(t, stats.Misses, int64(0))

	cache.ResetStats()

	stats = cache.GetStats()
	assert.Equal(t, int64(0), stats.Hits)
	assert.Equal(t, int64(0), stats.Misses)
	assert.Equal(t, int64(0), stats.Evictions)
}
