# https://www.acmicpc.net/problem/1004
import sys

def check_planet_contains(x, y, planet_x, planet_y, r):
    dist = (x-planet_x)**2 + (y-planet_y)**2
    if dist > r*r:
        return False
    return True

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        count = 0
        start_x, start_y, end_x, end_y = list(map(int, sys.stdin.readline().split()))
        n = int(sys.stdin.readline())
        for j in range(n):
            planet_x, planet_y, r = list(map(int, sys.stdin.readline().split()))
            is_start_contained = check_planet_contains(start_x, start_y, planet_x, planet_y, r)
            is_end_contained = check_planet_contains(end_x, end_y, planet_x, planet_y, r)
            if is_end_contained and not is_start_contained or not is_end_contained and is_start_contained:
                count += 1
        print(count)