# https://www.acmicpc.net/problem/3085
import sys

def swap(i, j, k, l, bonbons):
    bonbons[i][j], bonbons[k][l] = bonbons[k][l], bonbons[i][j]

def check_max(max_count, bonbons):
    prev = ""
    for i in range(len(bonbons)):
        count = 0
        for j in range(len(bonbons)):
            cur = bonbons[i][j]
            if prev == cur:
                count += 1
                max_count = max(count, max_count)
            else:
                count = 1
            prev = cur
    prev = ""
    for i in range(len(bonbons)):
        count = 0
        for j in range(len(bonbons)):
            cur = bonbons[j][i]
            if prev == cur:
                count += 1
                max_count = max(count, max_count)
            else:
                count = 1
            prev = cur
    return max_count

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    bonbons = []
    for i in range(n):
        inputs = list(sys.stdin.readline().rstrip())
        bonbons.append(inputs)
    max_count = 0
    for i in range(n):
        for j in range(n):
            if j != n-1:
                swap(i, j, i, j+1, bonbons)
                max_count = check_max(max_count, bonbons)
                swap(i, j, i, j+1, bonbons)
            if i != n-1:
                swap(i, j, i+1, j, bonbons)
                max_count = check_max(max_count, bonbons)
                swap(i, j, i+1, j, bonbons)
    print(max_count)