# https://www.acmicpc.net/problem/1934
import sys

def get_gcd(first, second):
    if first < second:
        second, first = first, second
    while second != 0:
        first, second = second, first%second
    return first
    
def get_lcm(first, second):
    return first*second / get_gcd(first, second)

if __name__ == "__main__":
    t = int(sys.stdin.readline())

    for i in range(t):
        a, b = list(map(int, sys.stdin.readline().split()))
        print(int(get_lcm(a, b)))

