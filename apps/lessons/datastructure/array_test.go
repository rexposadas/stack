package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// Basic tests for array. The point, being, creating and working with arrays.
func Test_Array(t *testing.T) {

	a := [3]int{1, 2}

	// Test that arrays are initialized to 0 when the type is an int.
	assert.Zero(t, a[2], "a[2] should be zero")
	assert.Equal(t, 1, a[0], "a[0] should be 1")
	assert.Equal(t, 2, a[1], "a[1] should be 2")
	assert.Lenf(t, a, 3, "a should have length of 3")

	// Arrays are mutable.
	for i, _ := range a {
		a[i] = a[i] * 10
	}

	assert.Equalf(t, 10, a[0], "a[0] should be 10")
	assert.Equalf(t, 20, a[1], "a[1] should be 20")
	assert.Equalf(t, 0, a[2], "a[2] should be 0")

	log.Printf("%+v", a)
}
