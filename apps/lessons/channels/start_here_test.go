package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// Simplest implementation of a channel.
func TestSimplest(t *testing.T) {
	m := make(chan int) // unbuffered, needs both a producer and receiver or there is a deadlock

	var wg sync.WaitGroup //

	wg.Add(1)
	go func() { // producer. in a separate routine add to the channel.
		defer wg.Done()
		m <- 1
	}()

	go func() { //
		wg.Wait()

		// if not closed, it does NOT produce a memory leak. the GC will take care of that.
		// This signals that there are no more content to be added to this channel.  This is how
		// the for-loop knows when to stop consuming the channel.
		close(m)
	}()

	// We can reach this part of the code because the functions above are running as separate routines.
	total := 0
	for v := range m { // loop until the channel is closed. see close(m) above.
		total += v
	}

	// We must have looped once.
	assert.Equal(t, 1, total)
}
