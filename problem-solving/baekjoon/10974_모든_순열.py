# https://www.acmicpc.net/problem/10974
import sys
from itertools import permutations

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    numbers = list(range(1, n+1))
    for i in permutations(numbers):
        for j in range(0, n):
            print("{} ".format(i[j]), end='')
        print()