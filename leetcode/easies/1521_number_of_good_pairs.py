
# https://leetcode.com/problems/number-of-good-pairs/description/

class Solution:
    def numIdenticalPairs(self, nums: list[int]) -> int:
        if len(nums) <= 1:
            return 0

        if len(nums) == 2:
            return 1

        pairs = 0
        for i in range(len(nums)-1):
            n = nums[i+1:]
            pairs+= n.count(nums[i]) # count how many times a number appears in a list.

        return pairs

