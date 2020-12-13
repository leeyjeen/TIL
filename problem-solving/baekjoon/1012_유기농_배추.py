# https://www.acmicpc.net/problem/1012
import sys

sys.setrecursionlimit(10000)

graph = []
visited = []

def dfs(row, col):
    global graph
    global visited
    visited[row][col] = True
    if row+1 < len(graph) and graph[row+1][col] == 1 and not visited[row+1][col]:
        dfs(row+1, col)
    if row-1 >= 0 and graph[row-1][col] == 1 and not visited[row-1][col]:
        dfs(row-1, col)
    if col+1 < len(graph[row]) and graph[row][col+1] == 1 and not visited[row][col+1]:
        dfs(row, col+1)
    if col-1 >= 0 and graph[row][col-1] == 1 and not visited[row][col-1]:
        dfs(row, col-1)

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        m, n, k = list(map(int, sys.stdin.readline().split()))
        
        graph = [[0 for j in range(m)] for l in range(n)]
        visited = [[False for j in range(m)] for l in range(n)]

        for j in range(k):
            x, y = list(map(int, sys.stdin.readline().split()))
            graph[y][x] = 1

        num_of_worm = 0
        for j in range(n):
            for l in range(m):
                if graph[j][l] == 1 and not visited[j][l]:
                    num_of_worm += 1
                    dfs(j, l)
        print(num_of_worm)