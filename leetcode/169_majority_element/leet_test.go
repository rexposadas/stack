package _69_majority_element

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/majority-element/
func majorityElement(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	targetSize := (len(nums) / 2) + 1
	var counter int
	var lastNum int
	for i, n := range nums {
		if i == 0 {
			lastNum = n
			counter++
			continue
		}

		// We found the majority element.
		if counter >= targetSize {
			return lastNum
		}

		if n == lastNum {
			counter++
		} else {
			lastNum = n
			counter = 1
		}
	}

	return lastNum
}

func Test_MajorityElement(t *testing.T) {

	cases := []struct {
		Input []int
		Out   int
	}{
		{Input: []int{3, 2, 3}, Out: 3},
		{Input: []int{2, 2, 1, 1, 1, 2, 2}, Out: 2},
	}

	for _, c := range cases {
		got := majorityElement(c.Input)
		assert.EqualValues(t, c.Out, got, "Input: %v", c.Input)
	}
}
