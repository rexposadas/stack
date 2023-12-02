
# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.

# desc: https://adventofcode.com/2023/day/1#part1
# Input https://adventofcode.com/2023/day/1/input
def process_line(line):
    # Initialize variables to store the first and last single-digit numbers
    first_digit = None
    last_digit = None

    # Iterate through each character in the line
    for char in line:
        if char.isdigit():  # Check if the character is a digit
            if first_digit is None:
                first_digit = char
            last_digit = char

    result = 0
    if first_digit is not None and last_digit is not None:
        # Return the sum of the first and last digits
        result =  int(first_digit + last_digit)
    elif first_digit is not None:
        # Return twice the value of the single digit if only one is found
        result =  int(first_digit + first_digit)


    print (f"Line: {line} Result: {result}")

    return result


def day1(name):
    file_path = name

    result = 0
    with open(file_path, "r") as file:
        for line in file:
            result += process_line(line.strip())

    return result


if __name__ == '__main__':
    result = day1("input.txt")
    print(result)