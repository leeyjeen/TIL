# https://www.acmicpc.net/problem/1010
import sys

dp = [[0]*31 for _ in range(31)]

def combi(m, n):
    if m == n or n == 0:
        return 1
    if dp[m][n] != 0:
        return dp[m][n]
    dp[m][n] = combi(m-1, n) + combi(m-1, n-1)
    return dp[m][n]

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        n, m = list(map(int, sys.stdin.readline().split()))
        print(combi(m, n))
