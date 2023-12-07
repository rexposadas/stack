package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

// Why use error groups?
// We have a set of routines ( like a file transfer to multiple locations )
// and we want to know if any of the routines failed.
func TestNoError(t *testing.T) {

	// designed to wait for a group of routines
	// note that we don't have to use a sync.WaitGroup
	eg := errgroup.Group{}

	eg.Go(func() error {
		return nil
	})

	eg.Go(func() error {
		return nil
	})

	// with for all routines to return. grab the first error.
	err := eg.Wait()

	// none of our routines returned an error
	assert.NoError(t, err)
}

func TestWithError(t *testing.T) {

	eg := errgroup.Group{}

	eg.Go(func() error {
		return nil
	})

	eg.Go(func() error {
		return fmt.Errorf("found")
	})

	err := eg.Wait()

	// one of the routines returned an error.
	assert.Error(t, err)
}
