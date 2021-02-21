# https://www.acmicpc.net/problem/4796
import sys

if __name__ == "__main__":
    i = 1
    while True:
        l, p, v = list(map(int, sys.stdin.readline().split()))
        if l == 0 and p == 0 and v == 0:
            break
        print("Case {}: {}".format(i, (v//p)*l + min(v%p, l)))
        i += 1