# https://www.acmicpc.net/problem/11052
import sys

def get_max_price(p):
    n = len(p)
    dp = [0 for _ in range(n)]
    dp[1] = p[1]
    for i in range(2, n):
        dp[i] = p[i]
        for j in range(1, i):
            dp[i] = max(dp[i], dp[i-j]+p[j])
    return dp[n-1]

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    p = [0]
    p.extend(list(map(int, sys.stdin.readline().split())))
    print(get_max_price(p))