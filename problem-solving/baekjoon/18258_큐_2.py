# https://www.acmicpc.net/problem/18258
# 큐의 개념을 익히고 실습하는 문제. 연산 당 시간 복잡도가 O(1)이어야 한다는 점에 유의.
"""
deque 활용.
Deques support thread-safe, memory efficient appends and pops from either side of the deque 
with approximately the same O(1) performance in either direction.
"""
import sys
from collections import deque   

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    dq = deque()
    for i in range(n):
        inputs = sys.stdin.readline().split()
        command = inputs[0]
        num = 0
        if len(inputs) > 1:
            num = inputs[1]
        if command == "push":
            dq.append(num)
        elif command == "pop":
            if len(dq) > 0:
                print(dq.popleft())
            else:
                print(-1)
        elif command == "size":
            print(len(dq))
        elif command == "empty":
            if len(dq) == 0:
                print(1)
            else:
                print(0)
        elif command == "front":
            if len(dq) > 0:
                print(dq[0])
            else:
                print(-1)
        elif command == "back":
            if len(dq) > 0:
                print(dq[-1])
            else:
                print(-1)

"""
# 시간초과 발생한 코드
# (list slicing 시간복잡도=O(n+k))
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    queue = []
    for i in range(n):
        inputs = sys.stdin.readline().split()
        command = inputs[0]
        num = 0
        if len(inputs) > 1:
            num = inputs[1]
        if command == "push":
            queue.append(num)   # 시간복잡도: O(1)
        elif command == "pop":
            if len(queue) > 1:
                print(queue[0])
                queue = queue[1:]   # 시간복잡도: O(n+k)
            elif len(queue) == 1:
                print(queue[0])
                queue = []
            else:
                print(-1)
        elif command == "size":
            print(len(queue))
        elif command == "empty":
            if len(queue) == 0:
                print(1)
            else:
                print(0)
        elif command == "front":
            if len(queue) > 0:
                print(queue[0])
            else:
                print(-1)
        elif command == "back":
            if len(queue) > 0:
                print(queue[len(queue)-1])
            else:
                print(-1)
"""