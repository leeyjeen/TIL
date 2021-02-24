# https://www.acmicpc.net/problem/13699
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    dp = [0] * (n+1)
    dp[0] = 1
    for i in range(1, n+1):
        for j in range(0, i):
            dp[i] += dp[j] * dp[i-j-1]
    print(dp[n])