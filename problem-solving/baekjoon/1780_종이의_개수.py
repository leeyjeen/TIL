# https://www.acmicpc.net/problem/1780
# 쿼드트리와 비슷한데 4개 대신 9개로 나누는 문제
import sys

paper = []
minus_count = 0
zero_count = 0
plus_count = 0

def check_paper(row, col, n):
    global paper, minus_count, zero_count, plus_count
    first = paper[row][col]
    is_complete = True
    for i in range(row, row+n):
        for j in range(col, col+n):
            if i == row and j == col:
                continue
            if paper[i][j] != first:
                is_complete = False
                break
        if not is_complete:
            break
    if is_complete:
        if first == -1:
            minus_count += 1
        elif first == 0:
            zero_count += 1
        elif first == 1:
            plus_count += 1
        return
    
    interval = n//3
    check_paper(row, col, interval)
    check_paper(row, col+interval, interval)
    check_paper(row, col+2*interval, interval)
    check_paper(row+interval, col, interval)
    check_paper(row+interval, col+interval, interval)
    check_paper(row+interval, col+2*interval, interval)
    check_paper(row+2*interval, col, interval)
    check_paper(row+2*interval, col+interval, interval)
    check_paper(row+2*interval, col+2*interval, interval)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    for i in range(n):
        paper.append(list(map(int, sys.stdin.readline().split())))

    check_paper(0, 0, n)
    print(minus_count)
    print(zero_count)
    print(plus_count)