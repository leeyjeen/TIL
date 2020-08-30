class StackNode:
    def __init__(self, data):
        self.data = data
        self.next = None
    
class Stack:
    def __init__(self):
        self.head = None
    
    def is_empty(self):
        return True if self.head is None else False
    
    def push(self, data):
        new_node = StackNode(data)
        new_node.next = self.head
        self.head = new_node
        
    def pop(self):
        if self.is_empty():
            return float("-inf")
        data = self.head.data
        self.head = self.head.next
        return data

    def peek(self):
        if self.is_empty():
            return float("-inf")
        return self.head.data

if __name__ == "__main__":
    stack = Stack()
    stack.push(10)
    stack.push(20)
    stack.push(30)
    stack.pop()
    print("Top element is %d " %(stack.peek()))