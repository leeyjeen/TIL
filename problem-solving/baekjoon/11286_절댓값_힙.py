# https://www.acmicpc.net/problem/11286
# 새로운 기준으로 뽑는 우선순위 큐를 만드는 문제
import sys, heapq

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    h = []

    for i in range(n):
        x = int(sys.stdin.readline())
        if x == 0:
            if len(h) > 0:
                print(heapq.heappop(h)[1])
            else:
                print(0)
        else:
            heapq.heappush(h, (abs(x), x))  # 튜플 활용