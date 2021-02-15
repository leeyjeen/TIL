# https://www.acmicpc.net/problem/20301
import sys

if __name__ == "__main__":
    n, k, m = list(map(int, sys.stdin.readline().split()))
    queue = [i+1 for i in range(n)]
    selected = 0
    removed = []
    turn_right = False
    while len(queue) > 0:
        if len(removed) % m == 0:
            turn_right = not turn_right
        if turn_right:
            selected = (selected + k - 1) % len(queue)
        else:
            selected = (selected - k) % len(queue)
            while selected < 0:
                selected += len(queue)
        removed.append(queue[selected])
        queue.pop(selected)
    for i in range(n):
        print(removed[i])