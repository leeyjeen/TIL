# https://www.acmicpc.net/problem/2309
import sys

if __name__ == "__main__":
    heights = []
    total = 0
    for i in range(0, 9):
        h = int(sys.stdin.readline())
        heights.append(h)
        total += h

    heights.sort()
    stop = False
    for i in range(0, 8):
        for j in range(i+1, 9):
            if heights[i]+heights[j] == total-100:
                for k in range(0, 9):
                    if k != i and k != j:
                        print(heights[k])
                stop = True
                break
        if stop:
            break