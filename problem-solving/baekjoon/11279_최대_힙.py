# https://www.acmicpc.net/problem/11279
# 최댓값을 빠르게 뽑는 자료구조를 배우는 문제
"""
heapq 모듈 활용:
heapq 모듈은 가장 작은 값을 루트로 갖는다.
그러므로 디폴트로 제공하는 자료구조는 최소 힙이다.
-> 최대 힙을 구현하려면, 값 입력 또는 출력시 마이너스를 곱하여 부호를 바꿔주면 된다.
"""
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
                print(-1*heapq.heappop(h))
        else:
            heapq.heappush(h, -1*x)
