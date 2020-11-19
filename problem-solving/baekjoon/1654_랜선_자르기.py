# https://www.acmicpc.net/problem/1654
# 흔히 parametric search라고도 부르는, 이분 탐색을 응용하여 최솟값이나 최댓값을 찾는 테크닉을 배우는 문제
import sys

def get_max_length(n, lengths):
    start = 1
    end = lengths[-1]
    mid = 0
    max_length = 0
    while start <= end:
        mid = (start + end) // 2
        total_count = 0
        for i in range(len(lengths)):
            total_count += lengths[i] // mid
        if total_count >= n:
            if max_length < mid:
                max_length = mid
                start = mid + 1
        elif total_count < n:
            end = mid - 1
    return max_length

if __name__ == "__main__":
    k, n = map(int, sys.stdin.readline().split())
    lengths = []
    for i in range(k):
        lengths.append(int(sys.stdin.readline()))
    lengths.sort()
    print(get_max_length(n, lengths))

    