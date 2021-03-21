# https://www.acmicpc.net/problem/11725
import sys

def bfs():
    queue = [1]
    parents = [0 for _ in range(len(tree))]
    while len(queue) > 0:
        parent = queue.pop(0)
        for v in tree[parent]:
            if not visited[v]:
                parents[v] = parent
                queue.append(v)
                visited[v] = True
    return parents

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    tree = [[] for _ in range(n+1)]
    visited = [False for _ in range(n+1)]
    for i in range(n-1):
        u, v = list(map(int, sys.stdin.readline().split()))
        tree[u].append(v)
        tree[v].append(u)
    parents = bfs()
    for i in range(2, n+1):
        print(parents[i])