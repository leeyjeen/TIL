# https://www.acmicpc.net/problem/10708
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    m = int(sys.stdin.readline())

    targets = list(map(int, sys.stdin.readline().split()))
    scores = [0 for i in range(n)]
    for i in range(m):
        target = targets[i]
        writes = list(map(int, sys.stdin.readline().split()))
        for j in range(n):
            if writes[j] == target:
                scores[j] += 1
            else:
                scores[target-1] += 1
    
    for i in range(n):
        print(scores[i])
