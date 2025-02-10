package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyInt int

func TestMain(t *testing.T) {
	assert.Equal(t, "blank", "blank")
}

func TestPut(t *testing.T) {
	c := NewLRUCache(2)
	c.Put("a", 1)
	assert.Equal(t, 1, len(c.cache))

	c.Put("b", 2)
	assert.Equal(t, 2, len(c.cache))

	assert.Equal(t, 1, c.get("a"))
	assert.Equal(t, 2, c.get("b"))

	// Start evicting since we can only fit 2 items.
	c.Put("c", 3) // evicts a ( least recently used )
	assert.Equal(t, 2, len(c.cache))
	assert.Equal(t, -1, c.get("a"))
}

func TestPutUpdate(t *testing.T) {
	c := NewLRUCache(2)
	c.Put("a", 1)
	assert.Equal(t, 1, len(c.cache))

	c.Put("a", 2)
	assert.Equal(t, 1, len(c.cache))
	assert.Equal(t, 2, c.get("a"))

	c.Put("b", 2)
	assert.Equal(t, 2, len(c.cache))
	assert.Equal(t, 2, c.get("b"))

	c.Put("b", 3)
	assert.Equal(t, 2, len(c.cache))
	assert.Equal(t, 3, c.get("b"))
}

func TestLRUCacheSimpleEvicted(t *testing.T) {
	c := NewLRUCache(2)
	c.Put("a", 1)
	c.Put("b", 2)
	assert.Equal(t, 1, c.get("a"))

	c.Put("c", 3) // evicts key b
	assert.Equal(t, -1, c.get("b"))
}

func TestEviction(t *testing.T) {
	c := NewLRUCache(3)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	assert.Equal(t, 3, len(c.cache))

	c.Put("d", 4) // evicts a
	assert.Equal(t, 3, len(c.cache))
	assert.Equal(t, -1, c.get("a"))
}

func TestLRUCacheSimpleThree(t *testing.T) {
	c := NewLRUCache(3)
	c.Put("a", 1)
	c.Put("b", 2)
	assert.Equal(t, 1, c.get("a"))

	c.Put("c", 3)
	assert.Equal(t, 2, c.get("b"))

	c.Put("d", 4) // evicts key 2
	assert.Equal(t, -1, c.get("a"))
	assert.Equal(t, 3, c.get("c"))
	assert.Equal(t, 4, c.get("d"))
}
