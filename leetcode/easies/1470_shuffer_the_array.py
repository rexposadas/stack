# https://leetcode.com/problems/shuffle-the-array/description/

class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        l = nums[:n]
        r = nums[n:]

        result = []
        for i in range(n
        ):
            result.append(l[i])
            result.append(r[i])

        return result
