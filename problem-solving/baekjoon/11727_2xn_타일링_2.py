# https://www.acmicpc.net/problem/11727
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    count = [0 for i in range(n+1)]
    for i in range(1, n+1):
        if i == 1:
            count[i] = 1
        elif i == 2:
            count[i] = 3
        else:
            count[i] = (count[i-1] + 2*count[i-2]) % 10007
    print(count[n])