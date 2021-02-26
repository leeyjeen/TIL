# https://www.acmicpc.net/problem/17212
import sys

dp = []
coins = [7, 5, 2, 1]

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    dp = [100000 for _ in range(n+1)]
    dp[0] = 0
    for i in range(1, n+1):
        for coin in coins:
            if i-coin >= 0 and dp[i-coin]+1 < dp[i]:
                dp[i] = dp[i-coin] + 1
    print(dp[n])
