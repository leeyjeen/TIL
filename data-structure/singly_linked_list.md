# Singly Linked List | 단일 연결 리스트
## 표현
연결 리스트는 연결 리스트의 첫 번째 노드를 가리키는 포인터에 의해 표현된다. 첫 번째 노드는 HEAD라고 부른다. 만약 연결 리스트가 비어 있다면, HEAD의 값은 NULL이다.  
리스트의 각 노드는 최소 두 부분으로 구성된다.
1) 데이터
2) 다음 노드를 가리키는 포인터

```python
# 노드 클래스
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# 연결 리스트 클래스
class LinkedList:
    def __init__(self):
        self.head = None

if __name__ == '__main__':
    # 빈 리스트 생성
    linked_list = LinkedList()

    # 세 개의 노드 생성
    linked_list.head = Node(1)
    second = Node(2)
    third = Node(3)
    
    # 첫 번째 노드를 두 번째 노드와 연결
    linked_list.head.next = second

    # 두 번째 노드를 세 번째 노드와 연결
    second.next = third
```

## 연결 리스트 순회
앞의 프로그램에서, 세 개의 노드와 간단한 연결 리스트를 생성하였다. 생성한 리스트를 순회하고 각 노드의 데이터를 출력해보자. 순회를 위하여, 주어진 리스트를 출력하는 범용 함수 printList()를 작성하도록 하자.

```python
# 연결 리스트의 순회를 위한 간단한 파이썬 프로그램
# 노드 클래스
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# 연결 리스트 클래스
class LinkedList:
    def __init__(self):
        self.head = None

    # 연결 리스트의 내용을 Head부터 시작해서 출력하는 함수
    def printList(self):
        temp = self.head
        while temp:
            print(temp.data)
            temp = temp.next

if __name__ == '__main__':
    # 빈 리스트 생성
    linked_list = LinkedList()

    # 세 개의 노드 생성
    linked_list.head = Node(1)
    second = Node(2)
    third = Node(3)
    
    # 첫 번째 노드를 두 번째 노드와 연결
    linked_list.head.next = second

    # 두 번째 노드를 세 번째 노드와 연결
    second.next = third

    linked_list.printList()
```
Output:
```bash
1
2
3
```

## 노드 삽입하기
노드를 추가하는 방식은 세 가지가 있다.
1) 연결 리스트의 맨 앞에 추가하기
2) 주어진 노드 다음에 추가하기
3) 연결 리스트의 마지막에 추가하기

## ```맨 앞에``` 노드 추가하기
- 새 노드는 항상 연결 리스트의 Head 전에 추가된다.  
- 새로 추가된 노드는 연결 리스트의 새로운 Head가 된다.  
(ex. 주어진 연결 리스트가 ```10->15->20->25```이고 ```5```를 맨 앞에 추가한다면, 연결 리스트는 ```5->10->15->20->25```가 된다.)
- ```push()```함수를 리스트의 맨 앞에 추가하는 함수라고 하자. ```push()```는 Head 포인터가 새로운 노드를 가리키도록 변경해야 하므로, Head 포인터를 가리키는 포인터를 입력 받아야 한다. (C의 경우)

![Alt text](./images/02.Linked_List_Insert.png?raw=true "Linked List Insert")

```python
# 연결 리스트 클래스 내부의 함수
# 시작 부분에 새로운 노드를 삽입하는 함수
def push(self, new_data):
    # 노드 할당 및 데이터 초기화
    new_node = Node(new_data)

    # 새 노드의 next가 현재의 head를 가리키도록 설정
    new_node.next = self.head

    # head가 새 노드를 가리키도록 설정
    self.head = new_node
```
```push()```의 시간 복잡도는 일정한 양의 작업을 하기 때문에 ```O(1)```이다.

## ```주어진 노드 뒤에``` 노드 추가하기
- 한 노드에 대한 포인터가 주어지고, 해당 노드 뒤에 새로운 노드를 추가한다면, 

![Alt text](./images/03.Linked_List_Insert2.png?raw=true "Linked List Insert")

```python
# 연결 리스트 클래스 내부의 함수
# 주어진 노드 뒤에 새 노드를 추가하는 함수
def insertAfter(self, prev_node, new_data):
    # 주어진 노드의 존재 여부 검사
    if prev_node is None:
        print("prev_node는 연결 리스트 내부에 존재하지 않습니다.")
        return
    
    # 새 노드 생성 및 데이터 초기화
    new_node = Node(new_data)

    # 새 노드가 이전 노드가 가리키던 것을 가리키도록 설정
    new_node.next = prev_node.next

    # 이전 노드가 새 노드를 가리키도록 설정
    prev_node.next = new_node
```
```insertAfter()```의 시간 복잡도는 일정한 양의 작업을 하기 때문에 ```O(1)```이다.

## ```맨 뒤에``` 노드 추가하기
새 노드는 항상 주어진 연결 리스트의 마지막 노드 뒤에 추가된다.  
예를 들어 ```5->10->15->20->25```의 연결리스트가 주어지고 ```30```이라는 아이템을 끝에 추가한다면, 연결 리스트는 ```5->10->15->20->25->30```이 된다.
연결 리스트는 일반적으로 그것의 head로 표현되기 떼문에, 리스트의 *끝까지 순회한 후* 마지막 노드를 새 노드로 변경하여야 한다.

![Alt text](./images/04.Linked_List_Insert3.png?raw=true "Linked List Insert")

```python
# 연결 리스트 클래스 내부의 함수
# 맨 끝에 노드를 추가하는 함수
def append(self, new_data):
    # 새 노드 생성 및 데이터 초기화
    new_node = Node(new_data)

    # 연결 리스트가 비어있는 경우, 새 노드를 head로 지정
    if self.head is None:
        self.head = new_node
        return
    
    # 마지막 노드까지 순회
    last = self.head
    while last.next:
        last = last.next
    
    # 마지막 노드 변경
    last.next = new_node
```
```append()```의 시간 복잡도는 head에서부터 끝까지 반복이 있기 때문에 연결 리스트 노드 개수가 n이라면 ```O(n)```이다. 이 메서드는 연결 리스트의 tail을 가리키는 추가 포인터를 이용하면 O(1)까지 최적화될 수 있다.

```python
if __name__ == "__main__":
    linked_list = LinkedList()
    
    linked_list.append(6) # 6->None
    linked_list.push(7) # 7->6->None
    linked_list.push(1) # 1->7->6->None
    linked_list.append(4) # 1->7->6->4->None
    linked_list.insertAfter(linked_list.head.next, 8) # 1->7->8->6->4->None

    linked_list.printList()
```
Output:
```bash
1
7
8
6
4
```

## 노드 삭제하기
key가 주어지면 연결 리스트에서 해당 key의 첫 번째 항목을 삭제하자.  
연결 리스트에서 노드를 삭제하기 위해서 다음 단계를 수행해야 한다.
1) 삭제할 노드의 이전 노드를 찾는다.
2) 이전 노드의 next를 변경한다.
3) 삭제할 노드의 메모리를 해제한다.

![Alt text](./images/05.Linked_List_Delete.png?raw=true "Linked List Delete")

```python
# 연결 리스트 클래스 내부의 함수
# 연결 리스트에서 key의 첫 번째 항목을 삭제하는 함수
def deleteNode(self, key):
    # head 노드 임시 보관
    temp = self.head

    # head 노드 자체가 삭제할 key를 가지고 있다면
    if temp is not None:
        if temp.data == key:
            self.head = temp.next
            temp = None
            return
    
    # 삭제할 key를 탐색한다.
    # prev.next를 변경하기 위해 이전 노드를 찾는다.
    while temp is not None:
        if temp.data == key:
            break
        prev = temp
        temp = temp.next
    
    # 연결 리스트에 key가 존재하지 않는다면
    if temp == None:
        return
    
    # 연결 리스트에서 노드 연결 해제
    prev.next = temp.next

    temp = None
```
```python
if __name__ == "__main__":
    linked_list = LinkedList()
    
    linked_list.append(6) # 6->None
    linked_list.push(7) # 7->6->None
    linked_list.push(1) # 1->7->6->None
    linked_list.append(4) # 1->7->6->4->None
    linked_list.deleteNode(6) # 1->7->4->None
    linked_list.insertAfter(linked_list.head.next, 8) # 1->7->8->4->None

    linked_list.printList()
```
Outptut:
```bash
1
7
8
4
```

## 주어진 위치의 노드 삭제하기
삭제할 노드가 root라면 단순하게 삭제한다. 중간 노드를 삭제하려면 삭제할 노드의 이전 노드를 가리키는 포인터를 가지고 있어야 한다. 따라서 position이 0이 아닌 경우 ```position-1```번의 반복을 수행 후 이전 노드를 가리키는 포인터를 얻어야 한다.

```python
# 연결 리스트 클래스 내부의 함수
# 리스트의 head로의 reference와 position이 주어졌을 때
# 주어진 position의 노드를 삭제하는 함수
def deleteNodeAt(self, position):
    # 연결 리스트가 비어 있는 경우
    if self.head == None:
        return
    
    # head 노드 임시 보관
    temp = self.head

    # 삭제할 노드가 head인 경우
    if position == 0:
        self.head = temp.next
        temp = None
        return

    # 삭제할 노드의 이전 노드 찾기
    for i in range(position-1):
        temp = temp.next
        if temp is None:
            break
    
    # position이 노드 개수보다 큰 경우
    if temp is None:
        return
    if temp.next is None:
        return
    
    # temp.next의 노드가 삭제할 노드인 경우
    # 삭제할 노드의 next를 가리키는 포인터를 보관
    next = temp.next.next

    # 연결 리스트에서 노드 연결 해제
    temp.next = next
```
```python   
if __name__ == "__main__":
    linked_list = LinkedList()
    
    linked_list.append(6) # 6->None
    linked_list.push(7) # 7->6->None
    linked_list.push(1) # 1->7->6->None
    linked_list.append(4) # 1->7->6->4->None
    linked_list.deleteNode(6) # 1->7->4->None
    linked_list.deleteNodeAt(1) # 1->4->None
    linked_list.insertAfter(linked_list.head.next, 8) # 1->4->8->None

    linked_list.printList()
```
Output:
```
1
4
8
```

## 연결 리스트 삭제하기
```python
# 연결 리스트 클래스 내부의 함수
# 연결 리스트를 순회하며 모든 노드를 삭제하는 함수
def deleteList(self):
    current = self.head
    while current:
        next_node = current.next
        del current.data
        current = next
```

## 연결 리스트 길이 찾기
주어진 단일 연결 리스트의 노드 개수를 세는 함수를 작성해보자.

![Alt text](./images/06.Linked_List_Length.png?raw=true "Linked List Length")

예를 들어 연결리스트 ```1->3->1->2->1```인 경우 5를 리턴해야 한다.

```python
# 연결 리스트 클래스 내부의 함수
# 연결 리스트의 전체 노드 개수를 세는 함수
def getCount(self):
    temp = self.head
    count = 0

    while temp:
        count += 1
        temp = temp.next
    return count
```

## 연결 리스트 요소 탐색하기
단일 연결 리스트에서 원하는 요소를 탐색하는 함수를 작성해보자.  
함수는 연결 리스트에 값이 존재할 경우 ```true```, 아닌 경우 ```false```를 반환한다.  

### 반복문으로 풀기
```python
# 연결 리스트 클래스 내부의 함수
# 연결 리스트에 특정 값이 존재하는지 검사하는 함수
def searchByIteration(self, x):
    current = self.head

    while current != None:
        if current.data == x:
            return True
        current = current.Next
    
    return False
```

### 재귀로 풀기
```python
# 연결 리스트 클래스 내부의 함수
# 연결 리스트에 특정 값이 존재하는지 검사하는 함수
def searchByRecursion(self, current, x):
    if not current:
        return False
    
    if current.data == x:
        return True

    return self.searchByRecursion(current.next, x)
```

## 연결 리스트의 N번째 노드 가져오기
```python
# 연결 리스트 클래스 내부의 함수
# 연결 리스트에서 주어진 인덱스 위치의 데이터를 반환하는 함수
def getNth(self, index):
    current = self.head
    count = 0
    
    while current:
        if count == index:
            return current.data
        count += 1
        current = current.next

    # 존재하지 않는 요소를 찾는 경우 assert fail..
    assert(False)
    return 0
```

## Reference
- https://www.geeksforgeeks.org/linked-list-set-1-introduction/
- https://www.geeksforgeeks.org/linked-list-set-2-inserting-a-node/
- https://www.geeksforgeeks.org/linked-list-set-3-deleting-node/?ref=lbp
- https://www.geeksforgeeks.org/delete-a-linked-list-node-at-a-given-position/?ref=lbp
- https://www.geeksforgeeks.org/write-a-function-to-delete-a-linked-list/?ref=lbp
- https://www.geeksforgeeks.org/find-length-of-a-linked-list-iterative-and-recursive/
- https://www.geeksforgeeks.org/search-an-element-in-a-linked-list-iterative-and-recursive/
- https://www.geeksforgeeks.org/write-a-function-to-get-nth-node-in-a-linked-list/