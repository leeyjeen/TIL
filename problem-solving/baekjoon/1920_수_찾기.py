# https://www.acmicpc.net/problem/1920
# 배열을 정렬한 후 이분 탐색으로 값을 찾아 봅시다.
import sys
import bisect

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    a = list(map(int, sys.stdin.readline().split()))
    a.sort()    # 정렬
    m = int(sys.stdin.readline())
    b = list(map(int, sys.stdin.readline().split()))

    for v in b:
        i = bisect.bisect_left(a, v)    # bisect 모듈 이용하여 이분탐색
        if i < len(a) and a[i] == v:
            print(1)
        else:
            print(0)
