# https://www.acmicpc.net/problem/1302
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    sold = {}
    for i in range(n):
        title = sys.stdin.readline().rstrip()
        if sold.get(title):
            sold[title] += 1
        else:
            sold[title] = 1

    counts = [(key, sold[key]) for _, key in enumerate(sold)]
    counts = sorted(counts, key=lambda x: (-x[1], x[0]))
    print(counts[0][0])