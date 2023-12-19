
# https://leetcode.com/problems/final-value-of-variable-after-performing-operations/description/
class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        v = 0
        for item in operations:
            if item == "--X" or item == "X--":
                v -= 1
            else:
                v += 1
        return v