# https://www.acmicpc.net/problem/2468
import sys

sys.setrecursionlimit(100000)

graph = []
visited = []

def dfs(row, col, h):
    global graph
    global visited
    
    visited[row][col] = True
    if row+1 < len(graph) and graph[row+1][col] >= h and not visited[row+1][col]:
        dfs(row+1, col, h)
    if row-1 >= 0 and graph[row-1][col] >= h and not visited[row-1][col]:
        dfs(row-1, col, h)
    if col+1 < len(graph) and graph[row][col+1] >= h and not visited[row][col+1]:
        dfs(row, col+1, h)
    if col-1 >= 0 and graph[row][col-1] >= h and not visited[row][col-1]:
        dfs(row, col-1, h)

def init_visited():
    global visited
    for i in range(len(visited)):
        for j in range(len(visited[i])):
            visited[i][j] = False

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    graph = [[0 for i in range(n)] for j in range(n)]
    visited = [[False for i in range(n)] for j in range(n)]
    for i in range(n):
        graph[i] = list(map(int, sys.stdin.readline().split()))

    max_num = 0
    for h in range(1, 101):
        num_of_safety = 0
        for i in range(n):
            for j in range(n):
                if graph[i][j] >= h and not visited[i][j]:
                    num_of_safety += 1
                    dfs(i, j, h)
        if num_of_safety > max_num:
            max_num = num_of_safety
        init_visited()
    print(max_num)