import unittest

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_sorted = sorted(s)
        t_sorted = sorted(t)

        return s_sorted == t_sorted





class TestSolution(unittest.TestCase):
    def setUp(self):
        # Create an instance of the Solution class before each test
        self.solution =  Solution()

    def test_example_1_true(self):
        # Test a list with duplicates
        s = "anagram"
        t = "nagaram"
        self.assertTrue(self.solution.isAnagram(s, t))

    def test_example_2_false(self):
        # Test a list with no duplicates
        s = "rat"
        t = "car"
        self.assertFalse(self.solution.isAnagram(s, t))

if __name__ == '__main__':
    unittest.main()
