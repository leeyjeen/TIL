# https://www.acmicpc.net/problem/1543
import sys

if __name__ == "__main__":
    document = sys.stdin.readline().rstrip()
    word = sys.stdin.readline().rstrip()
    print(document.count(word))