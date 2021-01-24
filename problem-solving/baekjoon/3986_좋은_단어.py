# https://www.acmicpc.net/problem/3986
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    count = 0
    for i in range(n):
        inputs = list(sys.stdin.readline().rstrip())
        stack = [inputs[0]]
        for j in range(1, len(inputs)):
            if len(stack) == 0:
                stack = [inputs[j]]
            elif stack[-1] == inputs[j]:
                stack.pop()
            else:
                stack.append(inputs[j])
        if len(stack) == 0:
            count += 1
    print(count)