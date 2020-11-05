# https://www.acmicpc.net/problem/9012
# 주어진 문자열이 올바른 괄호열인지 판단하는 문제
import sys

def is_vps(input_str):
    open_count, close_count = 0, 0
    for s in input_str:
        if s == "(":
            open_count += 1
        elif s == ")":
            close_count += 1
        if open_count < close_count:
            return "NO"
    if open_count != close_count:
        return "NO"
    return "YES"

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        print(is_vps(sys.stdin.readline()))