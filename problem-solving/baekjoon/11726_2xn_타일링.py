# https://www.acmicpc.net/problem/11726
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    count = [0 for i in range(n+1)]
    for i in range(1, n+1):
        if i < 3:
            count[i] = i
        else:
            count[i] = (count[i-1] + count[i-2]) % 10007
    print(count[n])