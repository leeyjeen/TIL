# https://www.acmicpc.net/problem/2294
import sys

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    coins = []
    for i in range(n):
        coins.append(int(sys.stdin.readline()))
    dp = [100000 for _ in range(k+1)]
    dp[0] = 0
    for i in range(1, len(dp)):
        for j in range(n):
            coin = coins[j]
            if i-coin >= 0:
                dp[i] = min(dp[i], dp[i-coin]+1)
    if dp[k] == 100000:
        print(-1)
    else:
        print(dp[k])