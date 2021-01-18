# https://www.acmicpc.net/problem/2635
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    max_count = 0
    answer = []
    for second in range(1, n+1):
        numbers = [n, second]
        i = 2
        while numbers[i-2]-numbers[i-1] >= 0:
            numbers.append(numbers[i-2]-numbers[i-1])
            if max_count < len(numbers):
                max_count = len(numbers)
                answer = numbers
            i += 1
    print(max_count)
    for i in answer:
        print("{} ".format(i), end='')
    print()