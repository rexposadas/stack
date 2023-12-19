# https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/

class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:

        out = []
        if len(nums) < 2:
            return out

        for i, n in enumerate(nums):
            smaller = 0
            for j, v in enumerate(nums):
                if i == j:
                    continue
                if v < n:
                    smaller += 1
            out.append(smaller)

        return out
