# https://www.acmicpc.net/submit/11478
import sys

if __name__ == "__main__":
    s = sys.stdin.readline().rstrip()
    strs = {}
    for i in range(len(s)):
        for j in range(1, len(s)-i+1):
            strs[s[i:i+j]] = True
    print(len(strs))