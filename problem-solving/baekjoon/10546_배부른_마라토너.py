# https://www.acmicpc.net/problem/10546
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    runners = {}
    for i in range(n):
        runner = sys.stdin.readline().rstrip()
        if runner in runners:
            runners[runner] += 1
        else:
            runners[runner] = 1
    for i in range(n-1):
        finished = sys.stdin.readline().rstrip()
        if runners[finished] == 1:
            runners.pop(finished)
        else:
            runners[finished] -= 1
    for i in runners.keys():
        print(i)