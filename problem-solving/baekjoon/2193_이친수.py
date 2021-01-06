# https://www.acmicpc.net/problem/2193
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    count = [0 for i in range(n+1)]
    for i in range(1, n+1):
        if i < 2:
            count[i] = 1
        else:
            count[i] = count[i-1] + count[i-2]
    print(count[n])