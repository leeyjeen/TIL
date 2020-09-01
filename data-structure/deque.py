class Deque:
    def __init__(self):
        self.deque = []

    def insertLast(self, data):
        self.deque.append(data)

    def insertFront(self, data):
        self.deque.insert(0, data)

    def deleteLast(self):
        if self.isEmpty():
            print("Deque is Empty")
        else:
            return self.deque.pop()

    def deleteFront(self):
        if self.isEmpty():
            print("Deque is Empty")
        else:
            return self.deque.pop(0)

    def isEmpty(self):
        if len(self.deque) == 0:
            return True
        return False

    def printDeque(self):
        for i in self.deque:
            print(i, end=" ")
        print()
        
if __name__ == "__main__":
    dq = Deque()
    dq.insertLast(3)
    dq.insertLast(5)
    dq.insertFront(1)
    dq.insertLast(7)
    dq.insertLast(9)
    dq.insertFront(-1)
    dq.printDeque()
    dq.deleteLast()
    dq.deleteFront()
    dq.printDeque()