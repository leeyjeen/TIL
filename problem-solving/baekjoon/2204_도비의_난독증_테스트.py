# https://www.acmicpc.net/problem/2204
import sys

if __name__ == "__main__":
    while True:
        n = int(sys.stdin.readline())
        if n == 0:
            break
        words = []
        for i in range(n):
            words.append(sys.stdin.readline().rstrip())
        print(sorted(words, key=lambda x: x.lower())[0])
