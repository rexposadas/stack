package main

import (
	"testing"

	"github.com/rexposadas/stack/lib/math"
	"github.com/stretchr/testify/assert"
)

// Notes: use our lib to demonstrate generics.
func TestSumGeneric(t *testing.T) {
	s := math.Sum(2, 4)
	assert.Equal(t, 6, s)

	f := math.Sum(2.2, 4.1, 1.0)
	assert.Equal(t, 7.3, f)
}
