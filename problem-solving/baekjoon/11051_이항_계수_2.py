# https://www.acmicpc.net/problem/11051
# 더 넓은 범위의 이항 계수를 동적 계획법으로 구하는 문제
import sys

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    # 팩토리얼 공식 대신 파스칼 삼각형을 이용하여 점화식 작성
    # nCk = nCn-k
    # nCk + nCk+1 = n+1Ck+1
    dp = [[0 for i in range(n+1)] for j in range(n+1)]
    dp[0][0] = dp[1][0] = dp[1][1] = 1
    for i in range(n+1):
        for j in range(n+1):
            if i==j or j==0:
                dp[i][j] = 1
            else:
                dp[i][j] = (dp[i-1][j-1] + dp[i-1][j])%10007
    print(dp[n][k])