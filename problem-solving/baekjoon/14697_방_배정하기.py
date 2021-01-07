# https://www.acmicpc.net/problem/14697
import sys

def check_rooms(a, b, c, n):
    for i in range(0, n+1):
        for j in range(0, n+1-i):
            for k in range(0, n+1-i-j):
                if a*i + b*j + c*k == n:
                    return 1
    return 0

if __name__ == "__main__":
    a, b, c, n = list(map(int, sys.stdin.readline().split()))
    print(check_rooms(a, b, c, n))