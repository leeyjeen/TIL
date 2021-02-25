# https://www.acmicpc.net/problem/8394
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    dp = [0, 1, 2]
    for i in range(3, n+1):
        dp.append((dp[i-1] + dp[i-2]) % 10)
    print(dp[n])