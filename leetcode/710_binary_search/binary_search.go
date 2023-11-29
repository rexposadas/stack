package leetcode

// https://leetcode.com/problems/binary-search/description/
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	for i, n := range nums {
		if n == target {
			return i
		}
	}

	return -1
}
