# https://programmers.co.kr/learn/courses/30/lessons/12969

import sys

n, m = map(int, sys.stdin.readline().split())
for i in range(0,m):
    for j in range(0,n):
        print('*', end='')
    print()