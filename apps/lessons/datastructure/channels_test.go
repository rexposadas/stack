package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test seen here: https://gobyexample.com/channels
func Test_SimpleChannel(t *testing.T) {

	m := make(chan string)
	go func() { m <- "ping" }()

	got := <-m // need a receiver for the one producer or the channel is locked
	assert.Equal(t, "ping", got)
}

// Buffered channels do not need a receiver to function properly.
func Test_Buffered(t *testing.T) {
	m := make(chan string, 2)

	m <- "one"
	m <- "two"

	assert.Equal(t, "one", <-m)
	assert.Equal(t, "two", <-m)
}

// Pass a pointer to the slice in order to
// modify it for the caller.
func growSlice(m chan bool, s *[]int) {
	*s = append(*s, 1)
	m <- true
}
func Test_Sync(t *testing.T) {
	s := []int{1}
	m := make(chan bool, 1)

	go growSlice(m, &s) // modifies the original slice
	<-m                 // Wait for the routine to modify the slice.

	// since we waited for the go routine to finish, s should
	// have been modified at this point.
	assert.Equal(t, 2, len(s))
}
