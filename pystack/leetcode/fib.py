

# Regular way
# O - nlogn
def fib(n):
    # base case
    if n == 0:
        return 0
    if n == 1:
        return 1

    # recursion
    return fib_map(n - 1) + fib_map(n - 2)

# Using a map so we don't process the same number repeatedly
# O - nlogn
m = {}
def fib_map(n):

    a = m.get(n)
    if a is not None:
        return a

    # base case
    if n == 0:
        m[n] = 0
        return 0

    if n== 1:
        m[n] = 1
        return 1


    # recursion 
    m[n] = fib_map(n - 1) + fib_map(n - 2)
    return m[n]


# constant memory
# O = n
def fib_const_mem(n):
    if n < 0:
        raise ValueError('n must be positive')
    if n == 0:
        return 0
    if n == 1:
        return 1

    prev, curr = 0, 1
    for _ in range(2, n+1):
        prev, curr = curr, prev + curr

    return curr


def test_fib():
    assert fib_map(0) == 0
    assert fib_map(1) == 1
    assert fib_map(2) == 1
    assert fib_map(3) == 2
    assert fib_map(4) == 3
    assert fib_map(5) == 5

def test_fib_map():
    assert fib_map(0) == 0
    assert fib_map(1) == 1
    assert fib_map(2) == 1
    assert fib_map(3) == 2
    assert fib_map(4) == 3
    assert fib_map(5) == 5

def test_fib_const_mem():
    assert fib_const_mem(0) == 0
    assert fib_const_mem(1) == 1
    assert fib_const_mem(2) == 1
    assert fib_const_mem(3) == 2
    assert fib_const_mem(4) == 3
    assert fib_const_mem(5) == 5


test_fib()
test_fib_map()
test_fib_const_mem()

print("all tests passed")  # Will print only if all tests passes
