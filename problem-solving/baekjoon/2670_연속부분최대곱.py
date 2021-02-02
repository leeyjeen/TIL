# https://www.acmicpc.net/problem/2670
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    real_numbers = []
    for i in range(n):
        real_numbers.append(float(sys.stdin.readline()))
    for i in range(1, n):
        mul = real_numbers[i-1] * real_numbers[i]
        real_numbers[i] = max(mul, real_numbers[i])
    answer = max(real_numbers)
    print("{:.3f}".format(answer))