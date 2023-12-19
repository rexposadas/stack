# https://leetcode.com/problems/goal-parser-interpretation/description/

class Solution:
    def interpret(self, command: str) -> str:

        final = ""
        total = ""

        for a in command:
            total += a
            if total == "G":
                final += "G"
                total = ""
                continue
            elif total == "()":
                final += "o"
                total = ""
                continue
            elif total == "(al)":
                final += "al"
                total = ""
                continue
            else:
                continue

        return final

