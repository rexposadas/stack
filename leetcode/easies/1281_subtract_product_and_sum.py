# https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/description/

class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        if n == 0:
            return 0

        sum = 0
        product = 1
        digits = [int(digit) for digit in str(n)]

        for a in digits:
            sum += a
            product *= a

        return product - sum