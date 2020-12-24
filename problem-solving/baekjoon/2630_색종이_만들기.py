# https://www.acmicpc.net/problem/2630
# 쿼드트리를 만드는 문제
import sys

paper = []

def check_part(row, col, n, blue, white):
    global paper
    first = paper[row][col]
    is_paper = True
    for i in range(row, row+n):
        for j in range(col, col+n):
            if i == row and j == col:
                continue
            if paper[i][j] != first:
                is_paper = False
                break
        if not is_paper:
            break

    if is_paper:
        if first == 1:
            blue += 1
        else:
            white += 1
        return blue, white
    
    interval = n//2
    blue, white = check_part(row, col, interval, blue, white)
    blue, white = check_part(row+interval, col, interval, blue, white)
    blue, white = check_part(row, col+interval, interval, blue, white)
    blue, white = check_part(row+interval, col+interval, interval, blue, white)
    return blue, white

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    for i in range(n):
        paper.append(list(map(int, sys.stdin.readline().split())))
    blue, white = check_part(0, 0, n, 0, 0)
    print(white)
    print(blue)