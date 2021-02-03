# https://www.acmicpc.net/problem/1822
import sys

if __name__ == "__main__":
    a_count, b_count = list(map(int, sys.stdin.readline().split()))
    a = set(map(int, sys.stdin.readline().split()))
    b = set(map(int, sys.stdin.readline().split()))
    diff = a.difference(b)
    result = sorted(list(diff))
    if len(result) == 0:
        print(0)
    else:
        print(len(result))
        for i in result:
            print("{} ".format(i), end='')
        print()