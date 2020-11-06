# https://www.acmicpc.net/problem/4949
# 주어진 문자열이 올바른 괄호열인지 판단하는 문제2
import sys

def is_balanced(input):
    stack = []
    for v in input:
        if v == "(" or v == "[":
            stack.append(v)
        elif v == ")":
            if len(stack) > 0 and stack[-1] == "(":
                stack = stack[:-1]
            else:
                return "no"
        elif v == "]":
            if len(stack) > 0 and stack[-1] == "[":
                stack = stack[:-1]
            else:
                return "no"
        else:
            continue
    if len(stack) > 0:
        return "no"
    return "yes"

if __name__ == "__main__":
    while True:
        input = sys.stdin.readline().rstrip('\n')
        if input == ".":
            break
        print(is_balanced(input))