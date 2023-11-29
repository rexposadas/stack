package __two_sum

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-sum/submissions/
func twoSum(nums []int, target int) []int {

	size := len(nums)

	for i := 0; i < size-1; i++ {
		for j := 1; j <= size-1; j++ {
			if i == j || i > j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func Test_TwoSum(t *testing.T) {

	cases := []struct {
		Nums   []int
		Target int
		Want   []int
	}{
		{Nums: []int{2, 7, 11, 15}, Target: 9, Want: []int{0, 1}},
		{Nums: []int{3, 2, 4}, Target: 6, Want: []int{1, 2}},
		{Nums: []int{0, 4, 3, 0}, Target: 0, Want: []int{0, 3}},
		{Nums: []int{-3, 4, 3, 90}, Target: 0, Want: []int{0, 2}},
	}

	for _, c := range cases {
		got := twoSum(c.Nums, c.Target)
		msg := fmt.Sprintf("expected %v, got %v, nums: %v, target: %d", c.Want, got, c.Nums, c.Target)
		assert.True(t, reflect.DeepEqual(c.Want, got), msg)
	}
}
