# https://www.acmicpc.net/problem/2217
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    ropes = []
    for i in range(n):
        ropes.append(int(sys.stdin.readline()))
    ropes.sort()
    max_weight = 0
    for i in reversed(range(len(ropes))):
        weight = ropes[i] * (len(ropes) - i)
        max_weight = max(max_weight, weight)
    print(max_weight)