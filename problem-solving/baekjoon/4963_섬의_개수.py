# https://www.acmicpc.net/problem/4963
import sys

sys.setrecursionlimit(2500) # python 재귀함수 호출 횟수는 1000회이므로 설정 변경

graph = []
visited = []

def dfs(h, w):
    global visited
    global graph
    visited[h][w] = True
    start_h, start_w, end_h, end_w = h, w, h, w
    if h-1 >= 0:
        start_h -= 1
    if h+1 < len(graph):
        end_h += 1
    if w+1 < len(graph[h]):
        end_w += 1
    if w-1 >= 0:
        start_w -= 1

    for i in range(start_h, end_h+1):
        for j in range(start_w, end_w+1):
            if graph[i][j] == 1 and not visited[i][j]:
                dfs(i, j)

if __name__ == "__main__":
    while True:
        w, h = list(map(int, sys.stdin.readline().split()))
        if w == 0 and h == 0:
            break

        graph = []
        visited = [[False for i in range(w)] for j in range(h)]
        for i in range(h):
            graph.append(list(map(int, sys.stdin.readline().split())))

        num_of_islands = 0
        for i in range(h):
            for j in range(w):
                if not visited[i][j] and graph[i][j] == 1:
                    num_of_islands += 1
                    dfs(i, j)
        print(num_of_islands)