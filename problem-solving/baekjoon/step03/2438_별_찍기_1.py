import sys

n = int(sys.stdin.readline())

for i in range(n):
    for j in range(i+1):
        print('*', end='')
    print()