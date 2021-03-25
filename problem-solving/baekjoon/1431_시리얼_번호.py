# https://www.acmicpc.net/problem/1431
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    serials = []
    for i in range(n):
        serials.append(sys.stdin.readline().rstrip())
    serials = sorted(serials, key=lambda s: (len(s), sum([int(i) for i in s if type(i)== int or i.isdigit()]), s))
    for i in serials:
        print(i)