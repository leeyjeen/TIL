# https://www.acmicpc.net/problem/11403
import sys

def floyd_warshall(graph):
    for i in range(len(graph)):
        for j in range(len(graph)):
            for k in range(len(graph)):
                if graph[j][i] != 0 and graph[i][k] != 0:
                    graph[j][k] = 1
    return graph

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    graph = []
    for i in range(n):
        graph.append(list(map(int, sys.stdin.readline().split())))

    result = floyd_warshall(graph)
    for i in range(n):
        for j in range(n):
            print("{} ".format(result[i][j]), end='')
        print()
