# https://leetcode.com/problems/build-array-from-permutation/description/

class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        ans = []

        for i in range (len(nums)):
            ans.insert(i, nums[nums[i]])

        return ans