# Doubly Linked List | 이중 연결 리스트
```이중 연결 리스트```는 일반적으로 단일 연결 리스트에 있는 ```next 포인터```, ```data```에 추가적으로 ```previous 포인터```라고 부르는 포인터를 포함한다. 

![Alt text](./images/07.Doubly_Linked_List.png?raw=true "Doubly Linked List")

## 표현
```python
# 이중 연결 리스트의 노드 클래스
class Node:
    def __init__(self, next=None, prev=None, data=None):
        self.next = next    # 다음 노드에 대한 참조
        self.prev = prev    # 이전 노드에 대한 참조
        self.data = data
```

## 단일 연결 리스트 대비 장점
1) 이중 연결 리스트는 양방향으로 순회할 수 있다.
2) 삭제할 노드에 대한 포인터가 주어진다면 이중 연결 리스트의 삭제 작업은 더 효율적이다.
3) 주어진 노드 앞에 빠르게 새로운 노드를 삽입할 수 있다.

*cf. 단일 연결 리스트에서 노드를 삭제하려면 이전 노드에 대한 포인터가 필요하다. 이전 노드를 얻기 위하여 리스트가 순회되어야 하는 경우가 있다. 이중 연결 리스트에서는 이전 포인터를 사용하여 이전 노드를 얻을 수 있다.*

## 단일 연결 리스트 대비 단점
1) 이중 연결 리스트의 모든 노드는 이전 포인터를 위한 추가 공간이 필요하다. 
2) 모든 연산은 수행 전 추가 포인터를 필요로 한다. 예를 들어 삽입의 경우, 이전 포인터와 다음 포인터를 모두 수정해야 한다. 즉, 아래에서 언급할 삽입을 위한 함수들은 이전 포인트를 설정하기 위해 1 또는 2개의 추가 단계가 필요하다.

## 노드 삽입하기
노드를 추가하는 방식은 네 가지가 있다.
1) 이중 연결 리스트의 맨 앞에 추가하기
2) 주어진 노드 다음에 추가하기
3) 이중 연결 리스트의 마지막에 추가하기 
4) 주어진 노드 이전에 추가하기

## ```맨 앞에``` 노드 추가하기
- 새 노드는 항상 주어진 연결 리스트의 Head 전에 추가된다. 
- 새로 추가된 노드는 이중 연결 리스트의 새로운 Head가 된다.  
(ex. 주어진 연결 리스트가 ```10<=>15<=>20<=>25```이고 ```5```를 맨 앞에 추가한다면, 연결 리스트는 ```5<=>10<=>15<=>20<=>25```가 된다.)
- ```push()```함수를 리스트의 맨 앞에 추가하는 함수라고 하자. ```push()```는 Head 포인터가 새로운 노드를 가리키도록 변경해야 하므로, Head 포인터를 가리키는 포인터를 입력 받아야 한다. (C의 경우)

![Alt text](./images/08.Doubly_Linked_List_Insert.png?raw=true "Doubly Linked List Insert")

```python
# 리스트의 맨 앞에 노드를 추가하는 함수
def push(self, new_data):
    # 노드 할당 및 데이터 초기화
    new_node = Node(data = new_data)

    # 새 노드의 next가 현재의 head를 가리키도록 설정
    new_node.next = self.head
    new_node.prev = None

    # head의 prev가 새 노드를 가리키도록 설정 (단일 연결 리스트에서 추가된 부분)
    if self.head is not None:
        self.head.prev = new_node
    
    # head가 새 노드를 가리키도록 설정
    self.head = new_node
```
단일 연결 리스트의 ```push()``` 프로세스에서 head의 이전 포인터를 변경하는 작업만 추가되었다.  

## ```주어진 노드 뒤에``` 노드 추가하기
- 이전 노드에 대한 포인터가 주어지고, 해당 노드 뒤에 새로운 노드를 한다. 

![Alt text](./images/09.Doubly_Linked_List_Insert2.png?raw=true "Doubly Linked List Insert")

```python
# 주어진 노드 뒤에 새 노드를 추가하는 함수
def insertAfter(self, prev_node, new_data):
    if prev_node is None:
        print("이중 연결 리스트에 해당 노드가 존재하지 않습니다.")
        return

    new_node = Node(data = new_data)
    new_node.next = prev_node.next
    prev_node.next = new_node
    new_node.prev = prev_node
    if new_node.next is not None:
        new_node.next.prev = new_node
```
단일 연결 리스트의 ```insertAfter()``` 함수에 새 노드의 이전 포인터를 변경하기 위한 단계와 새 노드의 다음 노드의 이전 포인터를 변경하기 위한 단계가 추가되었다.

## ```맨 뒤에``` 노드 추가하기
새 노드는 항상 주어진 연결 리스트의 마지막 노드 뒤에 추가된다.  
예를 들어 ```5<=>10<=>15<=>20<=>25```의 이중 연결 리스트가 주어지고 ```30```이라는 아이템을 끝에 추가한다면, 이중 연결 리스트는 ```5<=>10<=>15<=>20<=>25<=>30```이 된다.
연결 리스트는 일반적으로 그것의 head로 표현되기 떼문에, 리스트의 *끝까지 순회한 후* 마지막 노드를 새 노드로 변경하여야 한다.

![Alt text](./images/10.Doubly_Linked_List_Insert3.png?raw=true "Doubly Linked List Insert")


```python
# 이중 연결 리스트의 맨 뒤에 노드를 추가하는 함수
def append(self, new_data):
    new_node = Node(data = new_data)
    last = self.head

    new_node.next = None

    if self.head is None:
        new_node.prev = None
        self.head = new_node
        return

    while last.next is not None:
        last = last.next

    last.next = new_node
    new_node.prev = last
```

## ```주어진 노드 앞에``` 노드 추가하기

![Alt text](./images/11.Doubly_Linked_List_Insert4.png?raw=true "Doubly Linked List Insert")

```python
# 주어진 노드 앞에 새 노드를 추가하는 함수
def insertBefore(self, next_node, new_data):
    if next_node is None:
        print("이중 연결 리스트에 해당 노드가 존재하지 않습니다.")
        return

    new_node = Node(data = new_data)
    new_node.prev = next_node.prev
    next_node.prev = new_node
    new_node.next = next_node
    if new_node.prev is not None:
        new_node.prev.next = new_node
```

## 테스트
```python
# 이중 연결 리스트의 노드 클래스
class Node:
    def __init__(self, next=None, prev=None, data=None):
        self.next = next    # 다음 노드에 대한 참조
        self.prev = prev    # 이전 노드에 대한 참조
        self.data = data
        
# 이중 연결 리스트 클래스
class DoublyLinkedList:
    def __init__(self):
        self.head = None
        
    # 리스트의 맨 앞에 노드를 추가하는 함수
    def push(self, new_data):
        # 노드 할당 및 데이터 초기화
        new_node = Node(data = new_data)

        # 새 노드의 next가 현재의 head를 가리키도록 설정
        new_node.next = self.head
        new_node.prev = None

        # head의 prev가 새 노드를 가리키도록 설정 (단일 연결 리스트에서 추가된 부분)
        if self.head is not None:
            self.head.prev = new_node
        
        # head가 새 노드를 가리키도록 설정
        self.head = new_node
        
    # 주어진 노드 뒤에 새 노드를 추가하는 함수
    def insertAfter(self, prev_node, new_data):
        if prev_node is None:
            print("이중 연결 리스트에 해당 노드가 존재하지 않습니다.")
            return

        new_node = Node(data = new_data)
        new_node.next = prev_node.next
        prev_node.next = new_node
        new_node.prev = prev_node
        if new_node.next is not None:
            new_node.next.prev = new_node
            
    # 이중 연결 리스트의 맨 뒤에 노드를 추가하는 함수
    def append(self, new_data):
        new_node = Node(data = new_data)
        new_node.next = None

        if self.head is None:
            new_node.prev = None
            self.head = new_node
            return

        last = self.head
        while last.next is not None:
            last = last.next

        last.next = new_node
        new_node.prev = last
        
    # 주어진 노드 앞에 새 노드를 추가하는 함수
    def insertBefore(self, next_node, new_data):
        if next_node is None:
            print("이중 연결 리스트에 해당 노드가 존재하지 않습니다.")
            return

        new_node = Node(data = new_data)
        new_node.prev = next_node.prev
        next_node.prev = new_node
        new_node.next = next_node
        if new_node.prev is not None:
            new_node.prev.next = new_node
            
    # 연결 리스트의 내용을 주어진 노드에서부터 출력하는 함수
    def printList(self, node):
        print("정방향 순회")
        while node is not None:
            print("%d" % node.data)
            last = node
            node = node.next
            
        print("역방향 순회")
        while last is not None:
            print("%d" % last.data)
            last = last.prev
            
if __name__ == "__main__":
    dll = DoublyLinkedList()
    dll.append(6)   # 6
    dll.push(7)     # 7<=>6
    dll.push(1)     # 1<=>7<=>6
    dll.append(4)   # 1<=>7<=>6<=>4
    dll.insertAfter(dll.head.next, 8)   # 1<=>7<=>8<=>6<=>4
    dll.insertBefore(dll.head.next, 5)  # 1<=>5<=>7<=>8<=>6<=>4

    dll.printList(dll.head)
```
Output:
```bash
정방향 순회
1
5
7
8
6
4
역방향 순회
4
6
8
7
5
1
```

## Reference
- https://www.geeksforgeeks.org/doubly-linked-list/