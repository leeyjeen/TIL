# https://www.acmicpc.net/problem/1024
import sys

if __name__ == "__main__":
    n, l = list(map(int, sys.stdin.readline().split()))

    start = n//l - l//2
    total = 0
    is_valid = True
    while True:
        if start < 0:
            start += 1
        total = 0
        for i in range(start, start+l):
            total += i
        if total < n:
            start += 1
        elif total == n:
            if start < 0:
                is_valid = False
                break
            break
        elif total > n:
            l += 1
            if l > 100:
                is_valid = False
                break
            start = n//l - l//2
    if is_valid:
        for i in range(start, start+l):
            print("{} ".format(i), end='')
        print()
    else:
        print(-1)