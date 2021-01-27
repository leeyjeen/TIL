# https://www.acmicpc.net/problem/10799
import sys

def get_bar_count(inputs):
    stack = []
    prev = ""
    bar_count = 0
    for i in range(len(inputs)):
        cur = inputs[i]
        if cur == "(":
            stack.append(cur)
        elif cur == ")":
            stack.pop()
            if prev == "(":
                bar_count += len(stack)
            else:
                bar_count += 1
        prev = cur
    return bar_count

if __name__ == "__main__":
    inputs = list(sys.stdin.readline().rstrip())
    print(get_bar_count(inputs))