class CircularQueue():
    def __init__(self, size):
        self.size = size
        self.queue = [None for i in range(size)]
        self.front = -1
        self.rear = -1

    def enqueue(self, data):
        # 큐가 꽉 찬 경우
        if (self.rear + 1) % self.size == self.front:
            print("Queue is Full")
        elif self.front == -1:  # 큐가 비어 있는 경우
            self.front = 0
            self.rear = 0
            self.queue[self.rear] = data
        else:
            self.rear = (self.rear + 1) % self.size
            self.queue[self.rear] = data

    def dequeue(self):
        if self.front == -1: # 큐가 비어 있는 경우
            print("Queue is Empty")
        elif self.front == self.rear: # 큐에 하나의 데이터만 있는 경우
            temp = self.queue[self.front]
            self.front = -1
            self.rear = -1
            return temp
        else:
            temp = self.queue[self.front]
            self.front = (self.front + 1) % self.size
            return temp

    def print(self):
        if self.front == -1: # 큐가 비어 있는 경우
            print("Queue is Empty")
        elif self.rear >= self.front:
            for i in range(self.front, self.rear + 1):
                print(self.queue[i], end=" ")
            print()
        else:
            for i in range(self.front, self.size):
                print(self.queue[i], end=" ")
            for i in range(0, self.rear + 1):
                print(self.queue[i], end=" ")
            print()
        if (self.rear + 1) % self.size == self.front:
            print("Queue is Full")

if __name__ == "__main__":
    cq = CircularQueue(5)
    cq.enqueue(10)
    cq.enqueue(20)
    cq.enqueue(30)
    cq.enqueue(40)
    cq.enqueue(50)
    cq.print()
    cq.dequeue()
    cq.dequeue()
    cq.print()