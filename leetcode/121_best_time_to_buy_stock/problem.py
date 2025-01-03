import unittest

class Solution:
    def maxProfit(self, prices: list[int]) -> int:


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
