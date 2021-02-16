# https://www.acmicpc.net/problem/10994
import sys

stars = []

def draw_star(n, row, col):
    length = 4*n - 3
    if length == 1:
        stars[row][col] = "*"
        return
    for r in range(row, row+length):
        stars[r][col] = "*"
        stars[r][col+length-1] = "*"
    for c in range(col+1, col+length-1):
        stars[row][c] = "*"
        stars[row+length-1][c] = "*"
    draw_star(n-1, row+2, col+2)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    stars = [[" " for i in range(4*n-3)] for j in range(4*n-3)]
    draw_star(n, 0, 0)
    for i in stars:
        print("".join(i))