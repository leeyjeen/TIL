# https://www.acmicpc.net/problem/1021
# 덱을 활용하여 큐를 회전시키는 문제
import sys

class Queue:
     def __init__(self, cur_pos, init_pos):
        self.cur_pos = cur_pos
        self.init_pos = init_pos

def move_left(q, interval):
    result = []
    result.extend(q[interval:])
    result.extend(q[0:interval])
    return result

def move_right(q, interval):
    result = []
    result.extend(q[len(q)-interval:len(q)])
    result.extend(q[0:len(q)-interval])
    return result

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    
    q = []
    for i in range(n):
        q.append(Queue(i+1, i+1))
    
    select_pos = list(map(int, sys.stdin.readline().split()))
    count = 0
    while len(select_pos) > 0 and len(q) > 0:
        for i in range(len(q)):
            if q[i].init_pos == select_pos[0]:
                if i < len(q)-i:
                    count += i
                    q = move_left(q, i)
                else:
                    count += (len(q) - i)
                    q = move_right(q, len(q)-i)
                q = q[1:]
                select_pos = select_pos[1:]
                break
    print(count)