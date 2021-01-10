# https://www.acmicpc.net/problem/13301
import sys

def compute_sides(sides):
    sides[0] = 1
    sides[1] = 1
    for i in range(2, len(sides)):
        sides[i] = sides[i-1] + sides[i-2]
    return sides

def get_perimeter(sides):
    a = sides[len(sides)-1]
    b = sides[len(sides)-2]
    perimeter = a*4 + b*2
    return perimeter

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    if n == 1:
        print(4)
        sys.exit()
    elif n == 2:
        print(6)
        sys.exit()

    sides = [0 for i in range(n)]
    sides = compute_sides(sides)
    print(get_perimeter(sides))