# https://www.acmicpc.net/problem/10866
# 덱의 개념을 익히고 실습하는 문제
import sys

class Deque:
    def __init__(self):
        self.list = []
    
    def push_front(self, x):
        self.list.insert(0, x)
    
    def push_back(self, x):
        self.list.append(x)

    def pop_front(self):
        if self.is_empty():
            return -1
        return self.list.pop(0)

    def pop_back(self):
        if self.is_empty():
            return -1
        return self.list.pop(-1)

    def size(self):
        return len(self.list)

    def is_empty(self):
        if self.size() == 0:
            return True
        return False

    def front(self):
        if self.size() == 0:
            return -1
        return self.list[0]

    def back(self):
        if self.size() == 0:
            return -1
        return self.list[-1]

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    deque = Deque()
    for i in range(n):
        inputs = sys.stdin.readline().split()
        command = inputs[0]
        if command == "push_back":
            deque.push_back(inputs[1])
        elif command == "push_front":
            deque.push_front(inputs[1])
        elif command == "front":
            print(deque.front())
        elif command == "back":
            print(deque.back())
        elif command == "size":
            print(deque.size())
        elif command == "empty":
            if deque.is_empty():
                print(1)
            else:
                print(0)
        elif command == "pop_front":
            print(deque.pop_front())
        elif command == "pop_back":
            print(deque.pop_back())