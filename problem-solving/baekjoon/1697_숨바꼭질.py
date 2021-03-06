# https://www.acmicpc.net/problem/1697
import sys

def bfs(visited, start, end):
    queue = [start]
    while len(queue) > 0:
        cur_pos = queue.pop(0)
        if cur_pos == end:
            break
        if cur_pos-1 >= 0 and visited[cur_pos-1] == 0:
            queue.append(cur_pos-1)
            visited[cur_pos-1] += (visited[cur_pos]+1)
        if cur_pos+1 <= 100000 and visited[cur_pos+1] == 0:
            queue.append(cur_pos+1)
            visited[cur_pos+1] += (visited[cur_pos]+1)
        if cur_pos*2 <= 100000 and visited[cur_pos*2] == 0:
            queue.append(cur_pos*2)
            visited[cur_pos*2] += (visited[cur_pos]+1)

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    visited = [0 for _ in range(100001)]
    bfs(visited, n, k)
    print(visited[k])