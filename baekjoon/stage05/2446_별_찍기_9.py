import sys

n = int(sys.stdin.readline())

for i in range(0, n):
    for j in range(0, i):
        print(' ', end='')
    for j in range(0, 2*(n-i)-1):
        print('*', end='')
    print()
for i in range(0, n-1):
    for j in range(0, n-i-2):
        print(' ', end='')
    for j in range(0, 2*i+3):
        print('*', end='')
    print()