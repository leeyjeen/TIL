# https://www.acmicpc.net/problem/9465
import sys

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        n = int(sys.stdin.readline())
        scores = []
        dp = [[0]*n for _ in range(2)]
        for j in range(2):
            scores.append(list(map(int, sys.stdin.readline().split())))
        dp[0][0] = scores[0][0]
        dp[1][0] = scores[1][0]
        for j in range(1, n):
            if j == 1:
                dp[0][j] = dp[1][j-1] + scores[0][j]
                dp[1][j] = dp[0][j-1] + scores[1][j]
            else:
                dp[0][j] = max(dp[1][j-1], dp[1][j-2]) + scores[0][j]
                dp[1][j] = max(dp[0][j-1], dp[0][j-2]) + scores[1][j]
        print(max(dp[0][n-1], dp[1][n-1]))