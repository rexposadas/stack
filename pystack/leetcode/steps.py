class Solution:
    # O = 2^n
    # 2 recursive calls for every input  n ,
    # leading to an exponential growth in the number of calls:
    # T(n) = T(n-1) + T(n-2)

    # space complexity: n - the maximum depth of the recursion stack.
    def climbStairs(self, n: int) -> int:
        if n == 0 or n == 1:
            return 1

        return self.climbStairs(n - 1) + self.climbStairs(n - 2)

    def climbStairsLoop(self, n: int) -> int:
        if n == 0 or n == 1:
            return 1

        prev, curr = 1, 1
        for _ in range(2, n + 1):
            prev, curr = curr, prev + curr

        return curr


def test_stairs():
    print("running test using recursion")
    s = Solution()
    assert s.climbStairs(3) == 3
    assert s.climbStairs(2) == 2


def test_stairs_loop():
    print("running test using loop")

    s = Solution()
    assert s.climbStairsLoop(3) == 3
    assert s.climbStairsLoop(2) == 2


test_stairs()
test_stairs_loop()

print("all tests pass")  # will not print if any assert fails
