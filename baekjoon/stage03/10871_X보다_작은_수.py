import sys

n, x = map(int, sys.stdin.readline().split())
arr = list(map(int, sys.stdin.readline().split()))
for v in arr:
    if v < x:
        print(v, end=' ')
print()