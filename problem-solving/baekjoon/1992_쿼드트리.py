# https://www.acmicpc.net/problem/1992
# 쿼드트리를 문자열로 바꾸는 문제
import sys

image = []

def quad_tree(row, col, n):
    global image
    first = image[row][col]
    is_complete = True

    for i in range(row, row+n):
        for j in range(col, col+n):
            if i == row and j == col:
                continue
            if image[i][j] != first:
                is_complete = False
                break
        if not is_complete:
            break

    if is_complete:
        print(first, end='')
    else:
        print("(", end='')
        interval = n//2
        quad_tree(row, col, interval)
        quad_tree(row, col+interval, interval)
        quad_tree(row+interval, col, interval)
        quad_tree(row+interval, col+interval, interval)
        print(")", end='')

if __name__ == "__main__":
    n = int(sys.stdin.readline())

    for i in range(n):
        image.append(list(map(int, list(sys.stdin.readline().rstrip()))))
    
    quad_tree(0, 0, n)
    print()