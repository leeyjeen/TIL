# https://www.acmicpc.net/problem/2343
import sys

def blueray(lessons, m):
    low, high = 0, 0
    for lesson in lessons:
        high += lesson
        low = max(low, lesson)

    while low <= high:
        blueray_count = 0
        sum = 0
        mid = (low+high) // 2
        for i in range(len(lessons)):
            if sum+lessons[i] > mid:
                sum = 0
                blueray_count += 1
            sum += lessons[i]
        blueray_count += 1
        if blueray_count <= m:
            high = mid-1
        else:
            low = mid+1
    return low

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    lessons = list(map(int, sys.stdin.readline().split()))
    print(blueray(lessons, m))
