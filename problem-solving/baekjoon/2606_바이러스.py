# https://www.acmicpc.net/problem/2606
import sys

count = 0

def dfs(start):
    visited[start] = True
    global count
    count += 1

    for i in range(len(graph[start])):
        if graph[start][i] == 1 and not visited[i]:
            dfs(i)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    edge = int(sys.stdin.readline())

    graph = [[0 for i in range(n+1)] for j in range(n+1)]
    visited = [False for i in range(n+1)]

    for i in range(edge):
        computer1, computer2 = list(map(int, sys.stdin.readline().split()))
        graph[computer1][computer2] = 1
        graph[computer2][computer1] = 1

    dfs(1)
    print(count-1)