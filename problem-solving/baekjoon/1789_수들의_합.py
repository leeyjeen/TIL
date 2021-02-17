# https://www.acmicpc.net/problem/1789
import sys

if __name__ == "__main__":
    s = int(sys.stdin.readline())
    start, end = 1, s
    mid, answer = 0, 0
    while start <= end:
        mid = (start + end) // 2
        if mid*(mid + 1) // 2 <= s:
            answer = mid
            start = mid + 1
        else:
            end = mid - 1
    print(answer)