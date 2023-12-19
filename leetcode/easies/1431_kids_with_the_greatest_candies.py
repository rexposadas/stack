# https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/

class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:

        m = max(candies)
        r = []

        for c in candies:
            if c + extraCandies >= m:
                r.append(True)
            else:
                r.append(False)

        return r
