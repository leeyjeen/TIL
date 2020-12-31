# https://www.acmicpc.net/problem/11265
import sys

def floyd_warshall(graph):
    for i in range(len(graph)):
        for j in range(len(graph)):
            for k in range(len(graph)):
                if graph[j][i] + graph[i][k] < graph[j][k]:
                    graph[j][k] = graph[j][i] + graph[i][k]
    return graph

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    graph = []
    for i in range(n):
        graph.append(list(map(int, sys.stdin.readline().split())))

    graph = floyd_warshall(graph)

    for i in range(m):
        a, b, c = list(map(int, sys.stdin.readline().split()))
        if graph[a-1][b-1] <= c:
            print("Enjoy other party")
        else:
            print("Stay here")