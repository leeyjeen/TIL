# https://www.acmicpc.net/problem/11057
import sys

def get_asc_count(n):
    count = 0
    dp = [[0]*10 for _ in range(n+1)]
    for i in range(10):
        dp[1][i] = 1
    for i in range(1, n+1):
        for j in range(10):
            for k in range(j+1):
                dp[i][j] += dp[i-1][k]
                dp[i][j] %= 10007
    for i in range(10):
        count += dp[n][i]
    return count%10007

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    print(get_asc_count(n))