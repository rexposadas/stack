package _17_contains_duplicate

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {

	// Handle simple cases

	if len(nums) < 2 {
		return false
	}

	sort.Ints(nums)

	var last int
	for index, n := range nums {
		if index == 0 {
			last = n
			continue
		}

		if last == n {
			return true
		}

		last = n
	}

	return false
}

func Test_ContainsDuplicates(t *testing.T) {

	cases := []struct {
		Input []int
		Want  bool
	}{
		{Input: []int{1, 2, 3, 1}, Want: true},
		{Input: []int{1, 2, 3, 4}, Want: false},
		{Input: []int{}, Want: false},
		{Input: []int{1}, Want: false},
		{Input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, Want: true},
	}

	for _, c := range cases {
		got := containsDuplicate(c.Input)
		assert.EqualValues(t, c.Want, got, fmt.Sprintf("input: %v, expected: %v, got: %v", c.Input, c.Want, got))
	}
}
