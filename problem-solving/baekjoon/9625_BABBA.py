# https://www.acmicpc.net/problem/9625
import sys

def count_ab(k):
    a, b = 1, 0
    for i in range(k):
        prev_a, prev_b = a, b
        a += (prev_b - prev_a)
        b += prev_a
    return a, b

if __name__ == "__main__":
    k = int(sys.stdin.readline())
    a, b = count_ab(k)
    print("{} {}".format(a, b))