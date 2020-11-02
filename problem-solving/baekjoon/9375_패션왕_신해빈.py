
# https://www.acmicpc.net/problem/9375
# 경우의 수 연습문제
import sys

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        n = int(sys.stdin.readline())
        clothes = {}
        for j in range(n):
            name, kind = sys.stdin.readline().split()
            if clothes.get(kind):
                clothes[kind] += 1
            else:
                clothes[kind] = 1
        result = 1
        for v in clothes.values():
            result *= (v+1)
        result -= 1
        print(result)
