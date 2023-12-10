# see: https://leetcode.com/problems/valid-palindrome/
import unittest
class Solution:
    def isPalindrome(self, s: str) -> bool:
        new = ''
        for a in s:
            if a.isalpha() or a.isdigit():
                new += a.lower()

        # first semicolon ":" is the first index
        # second semicolon ":" is the last index
        # -1 means to iterate in reverse order
        return (new == new[::-1])

solution = Solution()

#  are palindromes
racecar = "racecar"
panama = "A man, a plan, a canal: Panama"
emtpy = " "

# not palindromes
car = "race a car"
boat = "sailboat"


class TestPalindrome(unittest.TestCase):
    def test_palindrome(self):
        self.assertEqual(solution.isPalindrome(racecar), True)
        self.assertEqual(solution.isPalindrome(panama), True)
        self.assertEqual(solution.isPalindrome(emtpy), True)

    def test_not_palindrome(self):
        self.assertEqual(solution.isPalindrome(car), False)
        self.assertEqual(solution.isPalindrome(boat), False)




