# Queue | 큐
```큐```는 **FIFO**(First In First Out, 선입선출) 순서를 따르는 **선형** 자료 구조이다.

큐에서는 주로 다음과 같은 4가지 기본 연산이 수행된다.
- Enqueue: 큐에 데이터 저장. 큐가 가득 찬 경우 overflow 조건.
- Dequeue: 큐에서 저장 순서가 가장 앞선 데이터를 삭제 및 반환. 스택이 비어 있는 경우 underflow 조건.
- Front: 큐의 맨 앞 항목 조회.
- Rear: 큐의 마지막 항목 조회.

## 큐의 응용
큐는 즉시 처리할 필요가 없지만 BFS와 같이 FIFO(선입선출) 순서로 처리되어야 할 때 사용된다. 큐의 이러한 속성은 다음과 같은 시나리오에서도 유용하다. 

1) 다수의 소비자 사이에 자원이 공유되는 경우. (ex. CPU 스케쥴링, 디스크 스케쥴링)
2) 두 프로세스 사이에 데이터가 비동기적으로 전송되는 경우 (ex. IO Buffers, pipes, file IO 등)

## 배열 기반 큐 구현 - Circular Queue (원형 큐)
```원형 큐```는 FIFO 원칙을 따르며 마지막 위치에서 다시 첫 번째 위치로 연결되며 원을 형성하는 선형 자료 구조이다. 링 버퍼(Ring Buffer)라고도 불린다.

일반 큐에서는, 큐가 가득 찰 때까지 삽입할 수 있다. 그러나 큐가 가득 차는 경우, 큐의 앞 부분에 공간이 남아 있어도 다음 항목을 삽입할 수 없다.

다음은 원형 큐의 연산들이다.
- Front: 큐의 맨 앞 항목 조회.
- Rear: 큐의 마지막 항목 조회.
- Enqueue: 원형 큐에 데이터 저장. 원형 큐에서 새 데이터는 항상 **Rear** 위치에 삽입된다.
1. 큐가 꽉 찼는지 검사한다. (rear == SIZE-1 && front == 0 || rear == front-1)
1. 꽉 찬 경우 Queue가 꽉 찼다고 알려준다. 그렇지 않은 경우 (rear == SIZE-1 && front != 0) 이면 ```rear=0``` 설정한다. 데이터를 삽입한다.  
- Dequeue: 원형 큐에서 저장 순서가 가장 앞선 데이터를 삭제 및 반환. 원형 큐에서 데이터는 항상 **Front** 위치에서 삭제된다. 
1. 큐가 비어 있는지 검사한다. (frotn == -1)
1. 비어 있는 경우 Queue가 비어 있다고 알려준다.
1. front == rear 인 경우 ```front = rear - 1``` 로 설정, 그렇지 않은 경우 front == size - 1 이면 ```front = 0```을 설정한다. 값을 반환한다.

```python
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
```
Output:
```bash
10 20 30 40 50 
Queue is Full
30 40 50
```
```enqueue```와 ```dequeue```의 시간 복잡도는 O(1)이다.

## 원형 큐의 응용
1. 메모리 관리: 일반 큐의 경우 사용되지 않는 메모리 위치를 원형 큐에서는 활용할 수 있다.
1. 트래픽 시스템: 컴퓨터 제어 교통 시스템에서 원형 큐는 설정된 시간에 따라 하나씩 반복적으로 신호등을 켜기 위해 원형 큐를 사용한다.
1. CPU 스케쥴링: 운영 체제는 종종 실행 준비가 되었거나 특정 이벤트가 발생하기를 기다리는 프로세스의 큐를 유지한다.

## 연결 리스트 기반 큐 구현
```python
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class Queue:
    def __init__(self):
        self.front = None
        self.rear = None
    
    def is_empty(self):
        return self.front == None

    def enqueue(self, data):
        temp = Node(data)
        if self.rear == None:
            self.front = self.rear = temp
            return
        self.rear.next = temp
        self.rear = temp

    def dequeue(self):
        if self.is_empty():
            return
        temp = self.front
        self.front = temp.next
        if self.front == None:
            self.rear = None
        return temp

if __name__ == "__main__":
    q = Queue()
    q.enqueue(10)
    q.enqueue(20)
    q.dequeue()
    q.dequeue()
    q.enqueue(30)
    q.enqueue(40)
    q.enqueue(50)
    q.dequeue()
    print("Queue Front " + str(q.front.data))
    print("Queue Rear " + str(q.rear.data))
```
Output:
```bash
Queue Front 40
Queue Rear 50
```
```enqueue```와 ```dequeue```의 시간 복잡도는 O(1)이다.

## References
- https://www.geeksforgeeks.org/queue-set-1introduction-and-array-implementation/
- https://www.geeksforgeeks.org/circular-queue-set-1-introduction-array-implementation/
- https://www.geeksforgeeks.org/queue-linked-list-implementation/