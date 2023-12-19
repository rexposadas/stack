
# https://leetcode.com/problems/jewels-and-stones/description/

class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        c = 0
        for j in jewels:
            a = stones.count(j)
            c += a
        return c


