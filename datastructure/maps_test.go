package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test working with basic maps.
func TestMaps(t *testing.T) {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	for k, v := range m {
		if k == "one" {
			assert.Equalf(t, 1, v, "key: %s, value: %d", k, v)
		}

		if k == "two" {
			assert.Equalf(t, 2, v, "key: %s, value: %d", k, v)
		}
	}
}

func TestMapInitialize(t *testing.T) {

	// These both initialize to an empty map.
	a := make(map[string]int)
	b := map[string]int{}

	assert.NotNil(t, a, "a should not be nil")
	assert.Equalf(t, len(a), 0, "a should have length of 0")

	assert.NotNil(t, b, "b should not be nil")
	assert.Equalf(t, len(b), 0, "b should have length of 0")
}
