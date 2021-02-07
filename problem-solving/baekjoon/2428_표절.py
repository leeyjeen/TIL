# https://www.acmicpc.net/problem/2428
import sys

def get_max_index(i, sizes):
    low, high = i+1, len(sizes)-1
    val = sizes[i]
    mid = 0
    max_index = i
    while low <= high:
        mid = (low + high) // 2
        if sizes[mid]*0.9 <= val:
            max_index = mid
            low = mid + 1
        else:
            high = mid - 1
    return max_index

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    sizes = list(map(float, sys.stdin.readline().split()))
    sizes.sort()
    total_count = 0
    for i in range(n-1):
        total_count += (get_max_index(i, sizes) - i)
    print(total_count)