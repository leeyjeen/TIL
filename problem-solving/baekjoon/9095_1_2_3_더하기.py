# https://www.acmicpc.net/problem/9095
import sys

if __name__ == "__main__":
    t = int(sys.stdin.readline())

    for i in range(t):
        n = int(sys.stdin.readline())
        count = [0 for j in range(n+1)]

        if n+1 > 1:
            count[1] = 1
        if n+1 > 2:
            count[2] = 2
        if n+1 > 3:
            count[3] = 4

        for i in range(4, n+1):
            count[i] = count[i-1] + count[i-2] + count[i-3]
        print(count[n])