# https://www.acmicpc.net/problem/2667
import sys

graph = []
visited = []
num_of_houses = 0

def dfs(row, col):
    global graph
    global visited
    global num_of_houses

    visited[row][col] = True
    num_of_houses += 1

    if row+1 < len(graph) and graph[row+1][col] == 1 and not visited[row+1][col]:
        dfs(row+1, col)
    if row-1 >= 0 and graph[row-1][col] == 1 and not visited[row-1][col]:
        dfs(row-1, col)
    if col+1 < len(graph) and graph[row][col+1] == 1 and not visited[row][col+1]:
        dfs(row, col+1)
    if col-1 >= 0 and graph[row][col-1] == 1 and not visited[row][col-1]:
        dfs(row, col-1)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    
    graph = [[0 for i in range(n)] for j in range(n)]
    visited = [[False for i in range(n)] for j in range(n)]

    for i in range(n):
        inputs = list(map(int, list(sys.stdin.readline().rstrip())))
        for j, val in enumerate(inputs):
            graph[i][j] = val
    numbers = []
    for i in range(n):
        for j in range(n):
            if graph[i][j] == 1 and not visited[i][j]:
                num_of_houses = 0
                dfs(i, j)
                numbers.append(num_of_houses)
    
    numbers.sort()
    print(len(numbers))
    for num in numbers:
        print(num)