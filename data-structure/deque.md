# Deque (Double-ended Queue) | 덱
```덱```은 양방향으로 삽입 및 삭제가 가능한 큐이다. 스택과 큐의 특성을 모두 갖는다. 

덱의 핵심 4가지 연산은 다음과 같다.
- insertFront(): 덱의 front에 삽입
- insertLast(): 덱의 rear에 삽입
- deleteFront(): 덱의 front에서 삭제
- deleteLast(): 덱의 rear에서 삭제

## 덱의 활용
덱은 스택과 큐 연산을 모두 지원하므로 둘 다로서 사용할 수 있다. 덱 자료 구조는 시계 방향 및 반시계 방향 회전을 특정 애플리케이션에서 유용하게 사용할 수 있는 O(1) 시간 내에 지원한다.

또한, 양방향으로 데이터의 삭제 또는 추가가 필요한 문제의 경우 덱을 이용하여 효율적으로 해결할 수 있다.

## 구현
```python
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
```
Output:
```bash
-1 1 3 5 7 9
1 3 5 7
```

## Reference
- https://www.geeksforgeeks.org/deque-set-1-introduction-applications/
- https://www.geeksforgeeks.org/implementation-deque-using-circular-array/