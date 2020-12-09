# https://www.acmicpc.net/problem/2644
import sys

graph = []
answer = 0

def dfs(person1, person2, prev, count):
    global answer
    if person1 == person2:
        answer = count
        return
    for i in range(len(graph)):
        if graph[person1][i] == 1 and i != prev:
            dfs(i, person2, person1, count+1)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    person1, person2 = list(map(int, sys.stdin.readline().split()))
    m = int(sys.stdin.readline())
    graph = [[0 for i in range(n+1)] for j in range(n+1)]

    for i in range(m):
        parent, child = list(map(int, sys.stdin.readline().split()))
        graph[parent][child] = 1
        graph[child][parent] = 1

    dfs(person1, person2, 0, 0)
    if answer == 0:
        answer = -1
    print(answer)