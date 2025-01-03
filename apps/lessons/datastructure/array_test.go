package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// Basic tests for arrays. The purpose is to create and work with arrays.
func Test_Array(t *testing.T) {

	a := [3]int{1, 2}

	// Verify that arrays are initialized to zero when the type is int.
	assert.Zero(t, a[2], "a[2] should be zero")
	assert.Equal(t, 1, a[0], "a[0] should be 1")
	assert.Equal(t, 2, a[1], "a[1] should be 2")
	assert.Len(t, a, 3, "a should have a length of 3")

	// Arrays can be modified.
	for i, _ := range a {
		a[i] = a[i] * 10
	}

	assert.Equalf(t, 10, a[0], "a[0] should be 10")
	assert.Equalf(t, 20, a[1], "a[1] should be 20")
	assert.Equalf(t, 0, a[2], "a[2] should be 0")

	log.Printf("%+v", a)
}
