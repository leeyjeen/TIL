# https://www.acmicpc.net/problem/16395
import sys

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    dp = [[0]*n for _ in range(n)]
    for i in range(n):
        for j in range(i+1):
            if j == 0 or j == i:
                dp[i][j] = 1
            else:
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
    print(dp[n-1][k-1])