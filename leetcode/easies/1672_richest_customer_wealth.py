# https://leetcode.com/problems/richest-customer-wealth/description/

class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        current_max = 0
        for a in accounts:
            m = 0
            for n in a:
                m += n
            if m > current_max:
                current_max = m
        return current_max
