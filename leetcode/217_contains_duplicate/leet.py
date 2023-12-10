import unittest

#  see: https://leetcode.com/problems/contains-duplicate/
class Solution:
    def containsDuplicate(self, nums: list[int]) -> bool:

        # use a set here because we don't care about order and we don't want duplicates.
        # it has the add method and the in operator is O(1).
        hashset = set()

        for n in nums:
            if n in hashset:
                return True
            hashset.add(n)
        return False


class TestSolution(unittest.TestCase):
    def setUp(self):
        # Create an instance of the Solution class before each test
        self.solution = Solution()

    def test_example_1_true(self):
        # Test a list with duplicates
        nums =  [1,2,3,1]
        self.assertTrue(self.solution.containsDuplicate(nums))

    def test_example_2_false(self):
        # Test a list with no duplicates
        nums = [1, 2, 3, 4]
        self.assertFalse(self.solution.containsDuplicate(nums))

    def test_empty_list(self):
        # Test an empty list
        nums = []
        self.assertFalse(self.solution.containsDuplicate(nums))

    def test_single_element(self):
        # Test a list with a single element
        nums = [1]
        self.assertFalse(self.solution.containsDuplicate(nums))

    def test_duplicate_at_end(self):
        # Test a list with a duplicate at the end
        nums = [1, 2, 3, 4, 4]
        self.assertTrue(self.solution.containsDuplicate(nums))

if __name__ == '__main__':
    unittest.main()
