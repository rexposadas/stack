#  https://leetcode.com/problems/number-of-employees-who-met-the-target/description/

class Solution:
    def numberOfEmployeesWhoMetTarget(self, hours: List[int], target: int) -> int:
        r = 0
        for item in hours:
            if item >= target:
                r += 1

        return r