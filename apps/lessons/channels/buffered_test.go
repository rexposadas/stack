package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffered(t *testing.T) {
	m := make(chan int, 2) // buffered.

	// The following lines cause deadlocks in unbuffered channels.
	// This works with buffered channels because we can hold the value in the buffer without a consumer.
	m <- 1   // buffer the number 1
	x := <-m // consume 1
	close(m)

	assert.Equal(t, 1, x)
}
