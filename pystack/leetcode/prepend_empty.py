# Given
# 1 - value
# 2 - array
# prepend value to the beginning of array
def prepend(a, value):
    raise NotImplementedError


a = [1, 2, 3]
assert prepend(a, 4) == [4, 1, 2, 3]
assert prepend(a, 1) == [4, 1, 2, 3]

print("OK: all tests passed")
