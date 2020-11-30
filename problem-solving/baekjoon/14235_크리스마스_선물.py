# https://www.acmicpc.net/problem/14235
import sys, heapq

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    h = []

    for i in range(n):
        inputs = list(map(int, sys.stdin.readline().split()))
        if inputs[0] == 0:
            if len(h) == 0:
                print(-1)
            else:
                print(-1*heapq.heappop(h))
        else:
            for j in range(1, inputs[0]+1):
                heapq.heappush(h, -1*inputs[j])