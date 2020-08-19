import sys

class Stack:
    def __init__(self):
        self.top = -1
        self.arr = []
    
    def get_size(self):
        return self.top + 1
    
    def is_empty(self):
        if self.top == -1:
            return True
        return False
    
    def is_full(self):
        if self.get_size() >= 10000:
            return True
        return False
        
    def push(self, n):
        if self.is_full():
            return
        self.top += 1
        self.arr.append(n)
        
    def pop(self):
        if self.is_empty():
            return -1
        result = self.arr[self.top]
        del self.arr[self.top]
        self.top -= 1
        return result
        
    def get_top(self):
        if self.is_empty():
            return -1
        return self.arr[self.top]

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    stack = Stack()
    for i in range(0, n):
        commands = sys.stdin.readline().split()
        if commands[0] == "push":
            stack.push(int(commands[1]))
        elif commands[0] == "pop":
            print(stack.pop())
        elif commands[0] == "size":
            print(stack.get_size())
        elif commands[0] == "empty":
            if stack.is_empty():
                print(1)
            else:
                print(0)
        elif commands[0] == "top":
            print(stack.get_top())