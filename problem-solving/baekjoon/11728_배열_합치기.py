# https://www.acmicpc.net/problem/11728
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    array = []
    array.extend(list(map(int, sys.stdin.readline().split())))
    array.extend(list(map(int, sys.stdin.readline().split())))
    array.sort()
    for i in range(len(array)-1):
        print(array[i], end=' ')
    print(array[-1])