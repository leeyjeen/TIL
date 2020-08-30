# Stack | 스택
```스택```은 **LIFO**(Last In First Out, 후입선출) 순서를 따르는 **선형** 자료 구조이다.

스택에서는 주로 다음과 같은 4가지 기본 연산이 수행된다.
- Push: 스택에 데이터 저장. 스택이 가득 찬 경우 overflow 조건.
- Pop: 스택에서 마지막으로 저장된 데이터 삭제 및 반환. 스택이 비어 있는 경우 underflow 조건.
- Peek: 스택에서 마지막으로 저장된 데이터 반환. 스택이 비어 있는 경우 underflow 조건.
- isEmpty: 스택이 빈 경우 True, 그렇지 않은 경우 False 반환.

## 스택 연산의 시간 복잡도
```push()```, ```pop()```, ```isEmpty()```, ```peek()```은 모두 O(1)의 시간 복잡도를 갖는다. 각 연산마다 아무런 반복문을 실행하지 않는다. 

## 스택의 응용
- 짝이 맞는 괄호 찾기 (ex. '{}()({{}})')
- Infix -> Postfix/Prefix 변환
- 에디터, 포토샵 등의 곳에서 Redo-Undo 기능
- 웹 브라우저에서 앞으로, 뒤로 가기 기능
- 하노이 타워, 트리 순회, 스톡 스팬 문제, 히스토그램 문제 등 많은 알고리즘에서 사용된다.

## 스택 구현
스택 구현은 2가지 방법으로 할 수 있다.
- 배열 기반 스택 구현
- 연결 리스트 기반 스택 구현

## 배열 기반 스택 구현
```python
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
```
Output:
```bash
30 popped from stack
```

## 연결 리스트 기반 스택 구현
새로운 노드를 머리에 추가하는 형태로 연결 리스트를 구현한다.
```python
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
```
Output:
```bash
Top element is 20
```


## Reference
- https://www.geeksforgeeks.org/stack-data-structure-introduction-program/