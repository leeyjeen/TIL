# https://www.acmicpc.net/problem/1057
import sys

if __name__ == "__main__":
    n, jimin, hansu = list(map(int, sys.stdin.readline().split()))
    round = 0
    while round < n:
        round += 1
        if jimin%2 == 0:
            jimin //= 2
        else:
            jimin = jimin//2 + 1
        if hansu%2 == 0:
            hansu //= 2
        else:
            hansu = hansu//2 + 1
        if jimin == hansu:
            print(round)
            sys.exit()
    print(-1)