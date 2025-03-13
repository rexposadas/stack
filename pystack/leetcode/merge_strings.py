# https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

# Given two strings word1 and word2. Merge the strings by
# adding letters in alternating order, starting with word1.
# If a string is longer than the other, append the additional
# letters onto the end of the merged string.

# Return the merged string.


class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        merged = []
        i, j = 0, 0

        # Merge characters alternately
        while i < len(word1) and j < len(word2):
            merged.append(word1[i])
            merged.append(word2[j])
            i += 1
            j += 1

        # Append the remaining characters from the longer string
        if i < len(word1):
            merged.append(word1[i:])
        if j < len(word2):
            merged.append(word2[j:])

        return "".join(merged)


def test_merge_strings():
    s = Solution()
    assert s.mergeAlternately("ab", "pqrs") == "apbqrs"
    assert s.mergeAlternately("abcd", "pq") == "apbqcd"


test_merge_strings()

print("all tests passed")  # will not print if any assert fails
