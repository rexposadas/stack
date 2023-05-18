package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func quickSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	return numbers
}

func Test_quicksort(t *testing.T) {
	a := struct {
		in       []int // input to the function
		expected []int // the expected sorted result
	}{}

	result := quickSort(a.in)
	assert.Equal(t, a.expected, result, "they should be equal")

}
