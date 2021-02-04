# https://www.acmicpc.net/problem/2910
import sys

if __name__ == "__main__":
    n, c = list(map(int, sys.stdin.readline().split()))
    frequency, order = {}, {}
    for i, num in enumerate(list(map(int, sys.stdin.readline().split()))):
        if frequency.get(num):
            frequency[num] += 1
        else:
            frequency[num] = 1
        if not order.get(num):
            order[num] = i + 1
    
    frequencies = []
    for i in frequency.keys():
        frequencies.append((i, frequency[i], order[i]))

    frequencies.sort(key=lambda a: (-a[1], a[2]))
    for i in frequencies:
        for j in range(i[1]):
            print("{} ".format(i[0]), end='')