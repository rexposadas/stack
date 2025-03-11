# https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

# Given two strings word1 and word2. Merge the strings by
# adding letters in alternating order, starting with word1.
# If a string is longer than the other, append the additional
# letters onto the end of the merged string.

# Return the merged string.

class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:

def test_merge_strings():
    s = Solution()
    assert s.mergeAlternately("ab","pqrs") == "apbqrs"
    assert s.mergeAlternately("abcd", "pq") == "apbqcd"


test_merge_strings()

print("all tests passed")  # will not print if any assert fails
