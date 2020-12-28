# https://www.acmicpc.net/problem/2740
# 행렬의 거듭제곱을 계산하기 전에 먼저 풀어야 할 문제
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    a = []
    for i in range(n):
        a.append(list(map(int, sys.stdin.readline().split())))

    m, k = list(map(int, sys.stdin.readline().split()))
    b = []
    for i in range(m):
        b.append(list(map(int, sys.stdin.readline().split())))

    result = [[0 for i in range(k)] for j in range(n)]
    for i in range(n):
        for j in range(k):
            for l in range(m):
                result[i][j] += a[i][l] * b[l][j]
    
    for i in range(n):
        for j in range(k):
            print("{} ".format(result[i][j]), end='')
        print()