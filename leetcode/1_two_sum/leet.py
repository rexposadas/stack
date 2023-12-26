# https://leetcode.com/problems/two-sum/description/

class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        prev = {}

        for i, n in enumerate(nums):
            diff = target - n
            if diff in prev: # look for the diff matching correct pair for n
                return [prev[diff], i]
            prev[n] = i # store the value and the index it was found
