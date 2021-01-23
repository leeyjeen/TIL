# https://www.acmicpc.net/problem/1620
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    dogam = {}
    for i in range(n):
        name = sys.stdin.readline().rstrip()
        dogam[name] = i+1
        dogam[str(i+1)] = name
    for i in range(m):
        question = sys.stdin.readline().rstrip()
        print(dogam[question])