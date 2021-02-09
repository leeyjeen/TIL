# https://www.acmicpc.net/problem/15815
import sys

def compute(operator, first, second):
    if operator == "+":
        return first + second
    elif operator == "-":
        return first - second
    elif operator == "*":
        return first * second
    else:
        return first // second

if __name__ == "__main__":
    formulas = list(sys.stdin.readline().rstrip())
    numbers = []
    while len(formulas) > 0:
        popped = formulas.pop(0)
        if popped == "+" or popped == "-" or popped == "*" or popped == "/":
            second = numbers.pop(-1)
            first = numbers.pop(-1)
            numbers.append(compute(popped, first, second))
        else:
            numbers.append(int(popped))
    print(numbers[0])