# https://leetcode.com/problems/running-sum-of-1d-array/description/

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        s = []
        current = 0
        for n in nums:
            total = n + current
            s.append(total)
            current = total

        return s
