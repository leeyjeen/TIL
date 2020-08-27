# Circular Linked List | 원형 연결 리스트
```원형 연결 리스트```는 모든 노드가 연결되어 원을 형성하는 연결 리스트이다. 끝에 NULL이 없다. 원형 연결 리스트는 단일 연결 리스트 또는 이중 연결 리스트 모두 될 수 있다.

![Alt text](./images/13.Circular_Linked_List.png?raw=true "Circular Linked List")

## 원형 연결 리스트의 장점
1) ```어떤 노드든 시작점```이 될 수 있다. 어느 지점에서 시작해도 전체 리스트를 순회할 수 있다. 처음 방문한 노드를 재방문하는 경우 정지하기만 하면 된다.
2) ```Queue``` 구현에 용이하다. 원형 연결 리스트를 이용하면 ```front```, ```rear``` 두 포인터를 유지할 필요가 없다. 마지막으로 삽입된 노드에 대한 포인터를 유지할 수 있으며 ```front```는 항상 마지막 노드의 next로 얻을 수 있다.
3) 원형 연결 리스트는 반복적으로 리스트를 순회하는 응용 프로그램에서 유용하다. 예를 들어, 여러 애플리케이션이 PC에서 실행 중일 떼, 운영체제가 실행 중인 애플리케이션들을 리스트에 올린 다음 이를 순환하며 각각 실행 시간을 한 슬라이스씩 부여한 다음 CPU를 다른 애플리케이션에 제공하는 동안 기다리게 하는 것이 일반적이다. 운영체제는 원형 리스트를 사용하면 리스트 끝에 도달했을 때 리스트의 맨 처음부터 순환할 수 있기 때문에 편리하다.
4) 원형 이중 연결 리스트는 ```Fibonacci Heap```과 같은 고급 자료 구조의 구현에 사용된다.

## 원형 연결 리스트 순회하기
기존의 연결 리스트에서는 Head 노드에서부터 리스트를 순회하다가 NULL에 도달하면 정지한다. 원형 연결 리스트에서는 *첫 번째 노드에 다시 도달*했을 때 정지한다.

## 코드
```python
# 원형 연결 리스트의 노드 클래스
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None
        
# 원형 연결 리스트 클래스
class CircularLinkedList:
    def __init__(self):
        self.head = None
        
    # 원형 연결 리스트의 맨 앞에 노드 삽입하는 함수
    def push(self, data):
        new_node = Node(data)
        temp = self.head
        
        new_node.next = self.head
        
        # 연결 리스트가 비어 있지 않은 경우 마지막 노드의 next 설정
        if self.head is not None:
            while temp.next != self.head:
                temp = temp.next
            temp.next = new_node
        else:
            new_node.next = new_node    # 연결 리스트가 비어 있는 경우 첫 번째 노드 설정
        
        self.head = new_node
        
    # 원형 연결 리스트의 노드들을 출력하는 함수    
    def printList(self):
        temp = self.head
        if temp is not None:
            while True:
                print("%d" % temp.data)
                temp = temp.next
                if temp == self.head:
                    break
                
if __name__ == "__main__":
    cll = CircularLinkedList()
    
    cll.push(12)    # 12
    cll.push(56)    # 56->12
    cll.push(2)     # 2->56->12
    cll.push(11)    # 11->2->56->12
    
    cll.printList()
```
Output:
```bash
11
2
56
12
```

## 노드 삽입하기
원형 연결 리스트를 구현할 때는 기존 연결 리스트와 다르게 리스트의 ```마지막 노드를 가리키는 추가 포인터```를 두는 것이 좋다. 마지막 노드를 가리키는 포인터를 가지고 있다면 마지막 노드의 ```next```를 이용하여 첫 번째 노드에도 접근이 가능하기 때문이다.

![Alt text](./images/14.Circular_Linked_List_Tail.png?raw=true "Circular Linked List Tail")

*cf. 왜 첫 번째 노드 대신 마지막 노드를 가리키는 포인터를 사용해야 하는가?
원형 연결 리스트의 맨 앞에 노드를 삽입하기 위해서 전체 리스트를 순회하여야 한다. 또한, 맨 끝에 삽입하기 위해서도 전체 리스트를 순회하여야 한다. 시작 포인터 대신 마지막 노드를 가리키는 포인터를 둔다면 두 경우 모두 전체 리스트를 순회할 필요가 없다. 따라서 앞 또는 끝에 삽입 작업이 리스트의 길이와 상관 없이 상수 시간이 걸린다.*

노드 추가는 네 가지 방법으로 이루어질 수 있다.
1) 비어 있는 리스트에 노드 추가하기
2) 리스트의 맨 앞에 노드 추가하기
3) 리스트의 맨 끝에 노드 추가하기
4) 노드 중간에 노드 추가하기

## 비어 있는 리스트에 노드 추가하기
to be continued..


## Referecnce
- https://www.geeksforgeeks.org/circular-linked-list/
- https://www.geeksforgeeks.org/circular-singly-linked-list-insertion/