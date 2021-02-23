# https://www.acmicpc.net/problem/9184
import sys

dp = [[[0 for i in range(21)] for j in range(21)] for k in range(21)]

def w(a, b, c):
    if a <= 0 or b <= 0 or c <= 0:
        return 1
    if a > 20 or b > 20 or c > 20:
        return w(20, 20, 20)
    if dp[a][b][c] != 0:
        return dp[a][b][c]
    if a < b and b < c:
        dp[a][b][c] = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
    else:
        dp[a][b][c] = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
    return dp[a][b][c]

if __name__ == "__main__":
    while True:
        a, b, c = list(map(int, sys.stdin.readline().split()))
        if a == -1 and b == -1 and c == -1:
            break
        print("w({}, {}, {}) = {}".format(a, b, c, w(a, b, c)))
