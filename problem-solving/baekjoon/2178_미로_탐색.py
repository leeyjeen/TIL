# https://www.acmicpc.net/problem/2178
import sys

count = []
graph = []
row_diff = [0, 0, -1, 1]
col_diff = [-1, 1, 0, 0]

def check(row, col):
    global graph
    if row >= len(graph) or row < 0 or col >= len(graph[0]) or col < 0:
        return False
    if graph[row][col] == 0 or count[row][col] != 0:
        return False
    return True

def bfs(row, col):
    global count
    global graph
    queue = [(row, col)]
    count[row][col] = 1
    while len(queue) > 0:
        r = queue[0][0]
        c = queue[0][1]
        if r == len(graph) and c == len(graph[0]):
            break
        queue.pop(0)

        for i in range(4):
            new_row = r + row_diff[i]
            new_col = c + col_diff[i]
            if check(new_row, new_col):
                queue.append((new_row, new_col))
                count[new_row][new_col] = count[r][c] + 1

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    count = [[0 for i in range(m)] for j in range(n)]
    graph = []

    for i in range(n):
        inputs = list(map(int, list(sys.stdin.readline().rstrip())))
        graph.append(inputs)

    bfs(0, 0)
    print(count[n-1][m-1])