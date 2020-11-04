# https://www.acmicpc.net/problem/2004
# nCm의 끝에 0이 얼마나 많이 오는지 구하는 문제
import sys

def get_two_count(n):
    i = 2
    two_count = 0
    while i <= n:
        two_count += n//i
        i*=2
    return two_count

def get_five_count(n):
    i = 5
    five_count = 0
    while i <= n:
        five_count += n//i
        i*=5
    return five_count

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    two_count = get_two_count(n) - get_two_count(n-m) - get_two_count(m)
    five_count = get_five_count(n) - get_five_count(n-m) - get_five_count(m)
    print(min(two_count, five_count))