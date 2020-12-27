# https://www.acmicpc.net/problem/1629
# 분할 정복으로 거듭제곱을 빠르게 계산하는 문제
import sys

sys.setrecursionlimit(10000)

def get_power(a, b, c):
    if b == 0:
        return 1
    elif b == 1:
        return a%c
    recur = get_power(a, b//2, c) % c
    if b % 2 == 0:
        return (recur % c) * (recur % c) % c
    return (recur % c) * (recur % c) * a % c

if __name__ == "__main__":
    a, b, c = list(map(int, sys.stdin.readline().split()))
    print(get_power(a, b, c))