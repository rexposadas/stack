import unittest

class Solution:
    def hammingWeight(self, n: int) -> int:
        res = 0

        # n starts as, for example, 00000000000000000000000000001011
        # until n is all zeros
        while n:
            n &= n - 1 # as long as there is a 1 bit, this run and removes it from n. n is changed here.
            res += 1    # and because we were able to set a 1 to a 0 (unset) we increment

        return res

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def run_test(self, test_cases):
        for n, expected in test_cases:
            with self.subTest(f"Testing with n={n}"):
                self.assertEqual(self.solution.hammingWeight(n), expected)

    def test_multiple_cases(self):
        test_cases = [
            (0, 0),
            (1, 1),
            (2, 1),
            (3, 2),
            (4, 1),
            (5, 2),
            (6, 2),
            (7, 3),
        ]

        self.run_test(test_cases)


if __name__ == '__main__':
    unittest.main()
