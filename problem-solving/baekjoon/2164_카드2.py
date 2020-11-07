# https://www.acmicpc.net/problem/2164
# 큐를 사용하여 동작을 구현하는 문제
import sys
from collections import deque

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    q = deque(maxlen=n)

    for i in range(n):
        q.append(i+1)
    
    discard = True
    while len(q) > 1:
        if discard:
            q.popleft()
        else:
            q.append(q.popleft())
        discard = not discard
    
    print(q.pop())