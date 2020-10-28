import sys

class Stack:
    def __init__(self):
        self.arr = []
        self.top = -1
      
    def get_size(self):
        return self.top + 1
    
    def is_empty(self):
        if self.top == -1:
            return True
        return False
        
    def push(self, n):
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
    k = int(sys.stdin.readline())
    stack = Stack()
    for i in range(0,k):
        num = int(sys.stdin.readline())
        if num == 0:
            stack.pop()
            continue
        stack.push(num)
    sum = 0
    while True:
        val = stack.pop()
        if val == -1:
            break
        sum += val
    print(sum)