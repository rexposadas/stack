package _28_summary_ranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {

	var answer []string

	var endingNumber int
	var startingNumber int

	// Take care of the simple cases.

	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	for i, num := range nums {

		// At the start, remember the first number
		if i == 0 {
			startingNumber = num
			endingNumber = num
			continue
		}

		// We reached the end of the slice
		if i == len(nums)-1 {
			if num == endingNumber+1 {
				answer = append(answer, makeStringArray(startingNumber, num))
				continue
			}
			// we have two ranges at this point.
			answer = append(answer, makeStringArray(startingNumber, endingNumber))
			answer = append(answer, makeStringArray(num, num))

			continue
		}

		if num == endingNumber+1 {
			endingNumber = num
			continue
		}

		// we just ended a range.
		answer = append(answer, makeStringArray(startingNumber, endingNumber))

		// start a new range
		startingNumber = num
		endingNumber = num
	}

	return answer
}

func makeStringArray(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	return fmt.Sprintf("%d->%d", start, end)
}

func Test_ValidParenthesis(t *testing.T) {

	cases := []struct {
		Input  []int
		Output []string
	}{
		{Input: []int{0, 1, 2, 4, 5, 7}, Output: []string{"0->2", "4->5", "7"}},
		{Input: []int{0, 2, 3, 4, 6, 8, 9}, Output: []string{"0", "2->4", "6", "8->9"}},
		{Input: []int{-1}, Output: []string{"-1"}},
	}

	for _, c := range cases {
		got := summaryRanges(c.Input)
		assert.EqualValues(t, c.Output, got, fmt.Sprintf("input: %v, expected: %v, got: %v", c.Input, c.Output, got))
	}
}
