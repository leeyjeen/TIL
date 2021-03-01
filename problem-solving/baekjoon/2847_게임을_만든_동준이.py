# https://www.acmicpc.net/problem/2847
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    scores = [int(sys.stdin.readline()) for _ in range(n)]
    result = 0 
    for i in range(n-2, -1, -1):
        if scores[i+1] <= scores[i]:
            diff = scores[i] - scores[i+1] + 1
            result += diff
            scores[i] -= diff
    print(result)