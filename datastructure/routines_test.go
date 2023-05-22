package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func Test_select(t *testing.T) {
	m := make(chan int)
	quit := make(chan int)

	go func() { // add to channel then exit
		m <- 1
	}()

	go func() { // add to channel then exit
		m <- 2
	}()

	go func() {
		// Let the other routines finish
		time.Sleep(1 * time.Second)
		quit <- 1
	}()

	// A select is used to listen to multiple channels
	for {
		select {
		case x := <-m:
			//t.Log("got: ", x)
			assert.Contains(t, []int{1, 2}, x)
		case <-quit:
			return
		}
	}
}

func Test_selectTimeOut(t *testing.T) {
	m := make(chan string)

	const d = 2 * time.Second
	stopper := time.After(d * time.Second)

	// Write to channel
	for i := 0; i < 4; i++ {
		go func() {
			m <- "good"
		}()
	}

	// a fun that should never finish since we will timeout before
	// this func finishes.
	go func() {
		time.Sleep((d + 1) * time.Second)
		// This shouldn't happen since the stopper should have stopped the program.
		m <- "bad"
	}()

	select {
	case x := <-m:
		// we shouldn't get a "bad" string since it happens after the stopper triggers.
		assert.Equal(t, "good", x)
	case <-stopper:
		log.Println("stop passed 2 seconds")
	}
}
