package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testSlice(t *testing.T, s []int) {
	assert.Equal(t, 2, len(s))
	assert.Equal(t, 1, s[0], "a[0] should be 1")
	assert.Equal(t, 2, s[1], "a[1] should be 2")
}

func TestSlice(t *testing.T) {
	// initialized. difference between slice and array, is that the length of the
	// array has to be declared.
	s := []int{1, 2}

	testSlice(t, s)
}

func TestSliceWithSize(t *testing.T) {
	s := make([]int, 2)
	s[0] = 1
	s[1] = 2
	testSlice(t, s)
}
