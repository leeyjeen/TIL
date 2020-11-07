# https://www.acmicpc.net/problem/11866
# 큐를 이용해 제거 과정을 구현하는 문제
import sys
from collections import deque

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))

    queue = deque(maxlen=n)
    for i in range(1, n+1):
        queue.append(i)
    
    result = []
    del_idx = 0
    while len(queue) > 0:
        del_idx = (del_idx+k-1)%len(queue)
        result.append(queue[del_idx])
        del queue[del_idx]
    
    for i, v in enumerate(result):
        if i == 0:
            print("<{}".format(v), end='')
        else:
            print(", {}".format(v), end='')
    print(">")