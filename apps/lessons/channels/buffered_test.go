package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffered(t *testing.T) {
	m := make(chan int, 2) // buffered.

	// Following lines causes deadlocks of unbuffered channels.
	// Works on buffered channels because we can hold the value in buffer without a consumer
	m <- 1   // buffer the number 1
	x := <-m // consume 1
	close(m)

	assert.Equal(t, 1, x)
}
