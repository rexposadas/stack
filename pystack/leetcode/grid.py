from typing import List

# https://leetcode.com/problems/number-of-islands/description/
# Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
#
# An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid:
            return 0

        count = 0
        for r in range(len(grid)):
            for c in range(len(grid[0])):
                if grid[r][c] == '1':
                    count += 1
                    self.dfs(grid, r, c)

        return count

    def dfs(self, grid, r, c):
        if r<0 or c<0 or r>=len(grid) or c>=len(grid[0]) or grid[r][c] != '1':
            return

        grid[r][c] = '0'
        self.dfs(grid, r-1, c)
        self.dfs(grid, r+1, c)
        self.dfs(grid, r, c-1)
        self.dfs(grid, r, c+1)



grid_1 = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
# expected Output: 1


grid_3 = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
# expected Output: 3

s = Solution()

assert s.numIslands(grid_1) == 1, f"{s.numIslands(grid_1)}"
assert s.numIslands(grid_3) == 3


print("all tests pass")



