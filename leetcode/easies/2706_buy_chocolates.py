# https://leetcode.com/problems/buy-two-chocolates/description/?envType=daily-question&envId=2023-12-20

class Solution:
    def buyChoco(self, prices: List[int], money: int) -> int:
        if len(prices) < 2:
            return money

        s = sorted(prices)
        d = money - (s[0] + s[1])
        if d >= 0:
            return d
        return money


