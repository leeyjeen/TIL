# https://www.acmicpc.net/problem/2947
import sys

if __name__ == "__main__":
    arr = list(map(int, sys.stdin.readline().split()))

    while arr[0] != 1 or arr[1] != 2 or arr[2] != 3 or arr[3] != 4 or arr[4] != 5:
        for i in range(4):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]
                print('{} {} {} {} {}'.format(arr[0], arr[1], arr[2], arr[3], arr[4]))
