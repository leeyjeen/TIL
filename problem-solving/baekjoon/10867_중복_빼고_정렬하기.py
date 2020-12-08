# https://www.acmicpc.net/problem/10867
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    numbers = set(map(int, sys.stdin.readline().split()))
    for i in sorted(numbers):
        print(i, end=" ")
    print()