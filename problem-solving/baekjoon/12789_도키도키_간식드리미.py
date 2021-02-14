# https://www.acmicpc.net/problem/12789
import sys

def get_result(students):
    stack = []
    line = []
    while len(students) > 0:
        if len(stack) == 0:
            stack.append(students.pop(0))
        else:
            if stack[-1] > students[0]:
                stack.append(students.pop(0))
            else:
                line.append(stack[-1])
                students.pop(0)
    while len(stack) > 0:
        line.append(stack[-1])
        stack.pop(-1)
    if line == sorted(line):
        return "Nice"
    return "Sad"

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    students = list(map(int, sys.stdin.readline().rstrip().split()))
    print(get_result(students))