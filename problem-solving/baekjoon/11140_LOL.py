# https://www.acmicpc.net/problem/11140
import sys
import string

def solve(input):
    if "lol" in input:
        return 0
    elif "lo" in input or "ol" in input or "ll" in input or any("l" + x + "l" in input for x in string.ascii_lowercase):
        return 1
    elif "l" in input or "o" in input:
        return 2
    return 3

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        input = sys.stdin.readline().rstrip()
        print(solve(input))