# https://www.acmicpc.net/problem/11653
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    factor = 2
    while factor <= n:
        if n%factor == 0:
            n /= factor
            print(factor)
        else:
            factor += 1