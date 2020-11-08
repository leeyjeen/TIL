# https://www.acmicpc.net/problem/1966
# 큐의 개념이 응용된 문제
import sys

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        n, m = list(map(int, sys.stdin.readline().split()))
        q = list(map(int, sys.stdin.readline().split()))
        queue = []
        for j, v in enumerate(q):
            queue.append((j, v))
        print_count = 0
        while len(queue) > 0:
            if queue[0][1] == max(q):
                print_count += 1
                if queue[0][0] == m:
                    print(print_count)
                    break
                else:
                    q.pop(0)
                    queue.pop(0)
            else:
                q.append(q.pop(0))
                queue.append(queue.pop(0))