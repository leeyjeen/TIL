# https://www.acmicpc.net/problem/11656
import sys

if __name__ == "__main__":
    s = sys.stdin.readline().rstrip()
    suffixes = []
    for i in range(len(s)):
        suffixes.append(s[i:])
    suffixes.sort()
    for suffix in suffixes:
        print(suffix)