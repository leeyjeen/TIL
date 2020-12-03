# https://www.acmicpc.net/problem/1260
import sys

graph = []
visited = []

def dfs(v):
    visited[v] = True
    print("{} ".format(v), end="")

    for i in range(len(graph[v])):
        if graph[v][i] == 1 and not visited[i]:
            dfs(i)

def bfs(v):
    visited[v] = True
    queue = [v]

    while (len(queue) != 0):
        front = queue.pop(0)
        print("{} ".format(front), end="")
        for i in range(len(graph[v])):
            if graph[front][i] == 1 and not visited[i]:
                visited[i] = True
                queue.append(i)

def reset_visited():
    for i in range(len(visited)):
        visited[i] = False

if __name__ == "__main__":
    n, m, v = map(int, sys.stdin.readline().split())

    graph = [[0 for i in range(n+1)] for j in range(n+1)]
    visited = [False for i in range(n+1)]

    for i in range(m):
        vertex1, vertex2 = map(int, sys.stdin.readline().split())
        graph[vertex1][vertex2] = 1
        graph[vertex2][vertex1] = 1
    
    dfs(v)
    print()

    reset_visited()
    bfs(v)
    print()