# https://www.acmicpc.net/problem/1764
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    dbj = {}
    for i in range(n):
        input = sys.stdin.readline().rstrip()
        if dbj.get(input):
            dbj[input] += 1
        else:
            dbj[input] = 1

    for i in range(m):
        input = sys.stdin.readline().rstrip()
        if dbj.get(input):
            dbj[input] += 1
        else:
            dbj[input] = 1

    dbj_result = []
    for (key, value) in dbj.items():
        if value == 2:
            dbj_result.append(key)

    dbj_result.sort()
    print(len(dbj_result))
    for i in dbj_result:
        print(i)