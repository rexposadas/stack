# https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/description/

class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        out = 0
        for s in sentences:
            parts = s.split()
            if len(parts) > out:
                out = len(parts)

        return out