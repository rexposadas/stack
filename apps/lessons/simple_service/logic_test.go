package main

// Demonstrates how to run tests.
// Use testify/assert because it's the best utility I've found for
// handling test cases.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Demonstrate how to run a test.
func TestAdd(t *testing.T) {
	result := Add(1, 2)
	assert.Equal(t, result, 3)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
}

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y int) {
		result := Add(x, y)
		assert.Equal(t, result, x+y)
	})
}
