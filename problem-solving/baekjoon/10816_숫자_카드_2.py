# https://www.acmicpc.net/problem/10816
# 이분 탐색으로 값의 개수를 찾아 봅시다.
import sys
import bisect

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    cards = list(map(int, sys.stdin.readline().split()))
    cards.sort()    # 정렬
    m = int(sys.stdin.readline())
    targets = list(map(int, sys.stdin.readline().split()))
    for t in targets:
        lower_bound = bisect.bisect_left(cards, t) # bisect 모듈 이용 lower bound 구하기
        upper_bound = bisect.bisect_right(cards, t) # bisect 모듈 이용 upper bound 구하기
        print(upper_bound - lower_bound, end=" ")
    print()