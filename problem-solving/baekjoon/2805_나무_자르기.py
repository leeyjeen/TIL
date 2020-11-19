# https://www.acmicpc.net/problem/2805
# 이분 탐색을 응용하여 최솟값이나 최댓값을 찾는 문제 2
import sys

def get_max_height(m, max_tree, trees):
    start, end, mid, max_height = 1, max_tree, 0, 0
    while start <= end:
        mid = (start + end) // 2
        total_height = 0
        for i in range(len(trees)):
            if trees[i] > mid:
                total_height += (trees[i] - mid)
        if total_height >= m:
            if max_height < mid:
                max_height = mid
                start = mid + 1
        else:
            end = mid - 1
    return max_height

if __name__ == "__main__":
    n, m = map(int, sys.stdin.readline().split())
    trees = list(map(int, sys.stdin.readline().split()))
    print(get_max_height(m, max(trees), trees))