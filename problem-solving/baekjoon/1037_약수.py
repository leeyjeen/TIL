# https://www.acmicpc.net/problem/1037
import sys

if __name__ == "__main__":
    count = int(sys.stdin.readline())
    factors = list(map(int, sys.stdin.readline().split()))
    print(min(factors)*max(factors))