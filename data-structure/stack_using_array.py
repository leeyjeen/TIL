# 스택이 비어 있는 경우 -infinite를 반환하기 위해 maxsize를 import한다.
from sys import maxsize
class Stack:
    def __init__(self):
        self.stack = []
    
    def push(self, item):
        self.stack.append(item)

    def pop(self):
        if self.is_empty():
            return -maxsize - 1
        return self.stack.pop()
    
    def peek(self):
        if self.is_empty():
            return -maxsize - 1
        return self.stack[-1]

    def is_empty(self):
        return len(self.stack) == 0

if __name__ == '__main__':
    stack = Stack()
    stack.push(10)
    stack.push(20)
    stack.push(30)
    print(str(stack.pop()) + " popped from stack")