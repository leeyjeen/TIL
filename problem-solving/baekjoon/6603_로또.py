# https://www.acmicpc.net/problem/6603
import sys
from itertools import combinations

if __name__ == "__main__":
    index = 0
    while True:
        inputs = list(map(int, sys.stdin.readline().split()))
        if inputs[0] == 0:
            break
        if index != 0:
            print()
        index += 1
        combi = list(combinations(inputs[1:], 6))
        for i, c in enumerate(combi):
            print(*list(c))
