# to do: a - b ( a and b are arrays)

# Given
# 1 - value
# 2 - array
# prepend value to the beginning of array
def prepend(a, value):
    set_a = set(a)
    if value in set_a:
        return a 

    a.insert(0, value)

    return a




a = [1, 2, 3]

assert prepend(a, 4) == [4,1,2,3]
assert prepend(a, 1) == [4,1,2,3]

print("OK: all tests passed")


