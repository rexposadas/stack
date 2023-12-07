package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	m := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(m)
	}()

	select {
	case x := <-m: // waits for values from routines
		log.Println(x)
	}

	// exits when channel closes and this test terminates.
}

func TestSelectWithDone(t *testing.T) {
	m := make(chan int)
	done := make(chan any)

	go func() { // periodically add to the m channel to be consumed by the select.

		// equivalent to doing a sleep like: time.Sleep(10 * time.Millisecond)
		ticker := time.NewTicker(10 * time.Millisecond)
		for { // keep looping. depend on done channel to exit the test.
			<-ticker.C
			m <- 1
		}
	}()

	// sleep, then end the test after a second. that's enough time to let the other
	// routine put values in the m channel
	go func() {
		ticker := time.NewTicker(1 * time.Second)

		<-ticker.C
		done <- true
	}()

	for { // need the infinite loop so that we can keep consuming the m channel
		select {
		case x := <-m:
			assert.Equal(t, 1, x)
		case <-done: // stop test after 1 second
			t.Log("done consumed")
			return // end this test
		}
	}
}
