# https://www.acmicpc.net/problem/14606
import sys

dp = []

def joy(n):
    if n == 1:
        return 0
    if n == 2:
        return 1
    dp[2] = 1
    dp[n] = (n//2)*(n-n//2) + joy(n//2) + joy(n-n//2)
    return dp[n]

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    dp = [0 for _ in range(n+1)]
    print(joy(n))