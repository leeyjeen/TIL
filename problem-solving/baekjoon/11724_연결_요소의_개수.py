# https://www.acmicpc.net/problem/11724
import sys

sys.setrecursionlimit(10000)

graph = []
visited = []

def dfs(start):
    global graph
    global visited
    visited[start] = True
    for i in range(len(graph)):
        if not visited[i] and graph[start][i] == 1:
            dfs(i)

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    graph = [[0 for i in range(n+1)] for j in range(n+1)]
    visited = [False for i in range(n+1)]

    for i in range(m):
        u, v = list(map(int, sys.stdin.readline().split()))
        graph[u][v] = 1
        graph[v][u] = 1

    num_of_connected = 0
    for i in range(n+1):
        if not visited[i]:
            num_of_connected += 1
            dfs(i)
    print(num_of_connected-1)
