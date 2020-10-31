# https://www.acmicpc.net/problem/11050
# N개의 물건 중 순서를 고려하지 않고 K개를 고르는 경우의 수, 이항 계수를 구하는 문제
import sys

def factorial(num):
    if num == 0 or num == 1:
        return 1
    return num*factorial(num-1)

if __name__ == "__main__":
    n, k = list(map(float, sys.stdin.readline().split()))
    print(int(factorial(n)/factorial(n-k)/factorial(k)))