# https://www.acmicpc.net/problem/13305
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    distances = list(map(int, sys.stdin.readline().split()))
    prices = list(map(int, sys.stdin.readline().split()))
    result = 0
    i = 0
    while i < n-1:
        cur = prices[i]
        dist = distances[i]
        for j in range(i+1, n-1):
            next = prices[j]
            if cur < next:
                dist += distances[j]
                i += 1
            else:
                break
        result += dist * cur
        i += 1
    print(result)