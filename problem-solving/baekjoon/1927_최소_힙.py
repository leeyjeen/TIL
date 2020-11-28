# https://www.acmicpc.net/problem/1927
# 최솟값을 빠르게 뽑는 문제
import sys, heapq

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    h = []

    for i in range(n):
        x = int(sys.stdin.readline())
        if x == 0:
            if len(h) == 0:
                print(0)
            else:
                print(heapq.heappop(h))
        else:
            heapq.heappush(h, x)