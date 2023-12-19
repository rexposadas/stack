
# https://leetcode.com/problems/find-words-containing-character/description/\

class Solution:
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:

        r = []

        for i, word in enumerate(words):
            if word.count(x) > 0:
                r.append(i)
        return r

