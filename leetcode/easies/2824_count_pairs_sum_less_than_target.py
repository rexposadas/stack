# https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/description/

class Solution:
    def countPairs(self, nums: List[int], target: int) -> int:
        if len(nums) < 2:
            return 0

        c = 0
        for i, n in enumerate(nums):
            for a in nums[i+1:]:
                if nums[i] + a < target:
                    c += 1

        return c
