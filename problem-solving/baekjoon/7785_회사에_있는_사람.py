# https://www.acmicpc.net/problem/7785
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())

    workers = {}
    for i in range(n):
        name, commute = sys.stdin.readline().split()
        if commute == "enter":
            workers[name] = 1
        elif commute == "leave":
            del workers[name]

    for i in sorted(list(workers.items()), reverse=True):
        print(i[0])