package _7_remove_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {

	size := len(nums)
	i := 0
	for i < size {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			size--
		} else {
			i++
		}
	}

	return len(nums)
}

func Test_RemoveElements(t *testing.T) {

	cases := []struct {
		Input []int
		Val   int
		Want  int
	}{
		{Input: []int{3, 2, 2, 3}, Val: 2, Want: 2},
		{Input: []int{0, 1, 2, 2, 3, 0, 4, 2}, Val: 2, Want: 5},
	}

	for _, c := range cases {
		got := removeElement(c.Input, c.Val)
		assert.EqualValues(t, c.Want, got, fmt.Sprintf("input: %v, val:%d, expected: %v, got: %v", c.Input, c.Val, c.Want, got))
	}
}
