# https://www.acmicpc.net/problem/2110
# 이분 탐색을 응용하여 최솟값이나 최댓값을 찾는 문제 3
import sys

def get_router_distance(x, c):
    start = 1
    end = x[-1]-x[0]
    dist = 0

    while start <= end:
        mid = (start + end) // 2
        prev = x[0]
        count = 1
        for i in range(len(x)):
            temp = x[i] - prev
            if mid <= temp :
                count += 1
                prev = x[i]

        if count >= c:
            dist = mid
            start = mid + 1
        else:
            end = mid - 1
    return dist


if __name__ == "__main__":
    n, c = map(int, sys.stdin.readline().split())
    x = []
    for i in range(n):
        x.append(int(sys.stdin.readline()))
    x.sort()
    print(get_router_distance(x, c))