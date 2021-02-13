# https://www.acmicpc.net/problem/6187
import sys
from itertools import combinations 

def get_heaviest(c, w):
    max_sum = 0
    for i in range(len(w), 0, -1):
        for v in combinations(w, i):
            sum = 0
            for j in range(0, len(v)):
                if sum+v[j] <= c:
                    sum += v[j]
                else:
                    break
            max_sum = max(max_sum, sum)
            if max_sum == c:
                return max_sum
    return max_sum

if __name__ == "__main__":
    c, n = map(int, sys.stdin.readline().split())
    w = [int(sys.stdin.readline()) for i in range(n)]
    print(get_heaviest(c, w))