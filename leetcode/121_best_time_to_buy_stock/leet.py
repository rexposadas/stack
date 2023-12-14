import unittest

class Solution:
    def maxProfit(self, prices: list[int]) -> int:

        if len(prices) <= 1:
            return 0

        if len(prices) == 2:
            if prices[1] > prices[0]:
                return prices[1] - prices[0]
            else:
                return 0

        max_profit = 0
        lowest_day = prices[0]
        for p in prices:
            if lowest_day > p:
                lowest_day = p
            max_profit = max(max_profit, p - lowest_day)

        return max_profit

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution =  Solution()

    def test_example_true(self):
        prices = [7,1,5,3,6,4]
        self.assertEqual(self.solution.maxProfit(prices), 5)

    def test_example_false(self):
        prices = [7,6,4,3,1]
        self.assertEqual(self.solution.maxProfit(prices), 0)


    def test_example_empty(self):
        prices = []
        self.assertEqual(self.solution.maxProfit(prices), 0)


    def test_example_one(self):
        prices = [1]
        self.assertEqual(self.solution.maxProfit(prices), 0)


    def test_example_two_zero(self):
        prices = [1, 1]
        self.assertEqual(self.solution.maxProfit(prices), 0)



    def test_example_two_2(self):
        prices = [1, 3]
        self.assertEqual(self.solution.maxProfit(prices), 2)



    def test_example_two_2(self):
        prices = [4, 3]
        self.assertEqual(self.solution.maxProfit(prices), 0)


if __name__ == '__main__':
    unittest.main()
