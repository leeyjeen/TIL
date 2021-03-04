# https://www.acmicpc.net/problem/7576
import sys

graph = []
queue = []
dcol = [-1, 0, 1, 0]
drow = [0, 1, 0, -1]

def bfs():
    while len(queue) > 0:
        row = queue[0][0]
        col = queue[0][1]
        queue.pop(0)
        for i in range(4):
            new_row = row + drow[i]
            new_col = col + dcol[i]
            if new_row < 0 or new_row > len(graph)-1 or new_col < 0 or new_col > len(graph[row])-1:
                continue
            if graph[new_row][new_col] != 0:
                continue
            graph[new_row][new_col] += (graph[row][col]+1)
            queue.append((new_row, new_col))

def get_time_to_grow_ripe(graph):
    total_time = 0
    for i in range(len(graph)):
        for j in range(len(graph[i])):
            if graph[i][j] == 0:
                return -1
            total_time = max(graph[i][j], total_time)
    return total_time-1

if __name__ == "__main__":
    m, n = list(map(int, sys.stdin.readline().split()))
    for i in range(n):
        graph.append(list(map(int, sys.stdin.readline().split())))
        for j in range(m):
            if graph[i][j] == 1:
                queue.append((i, j))
    bfs()
    print(get_time_to_grow_ripe(graph))