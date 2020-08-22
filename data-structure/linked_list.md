# Linked List | 연결 리스트

```연결 리스트```는 요소들이 ```인접한 메모리 위치에 저장되지 않는```(=배열과 차이점) ```선형 자료 구조```(=배열과 공통점)이다.  
연결 리스트의 요소들은 아래 이미지에 표시된 것처럼 ```포인터를 사용하여 연결```된다.   

![Alt text](./images/01.Linked_List.png?raw=true "Linked List")

연결 리스트는 ```노드```들로 구성되어 있으며, 각 노드는 ```데이터``` 필드와 리스트 내의 다음 노드에 대한 ```참조(포인터)```를 포함한다.

## 종류
- Singly Linked List | 단일 연결 리스트
- Doubly Linked List | 이중 연결 리스트
- Circular Linked List | 원형 연결 리스트

## 왜 연결리스트인가?
배열은 비슷한 형식의 선형 데이터를 저장하기 위해 사용되는데, 다음과 같은 한계를 가지고 있다.

1) 배열의 크기는 고정되어 있다.:  
그래서 배열 요소 개수의 상한을 미리 알고 있어야 한다.  
또한, 일반적으로 할당된 메모리는 사용량에 관계 없이 상한과 동일하다.  

2) 새로운 요소를 배열에 삽입할 때 비용이 크다.:  
새 요소를 위하여 공간이 생성되어야 하고 공간을 생성하기 위해서 기존 요소들을 이동 시켜야 한다.  

예를 들어, 한 시스템에서, 정렬된 id[] 배열 리스트를 유지하고 싶은 경우 
```
id[] = [1000, 1010, 1050, 2000, 2040]
```
여기에 새로운 ID 1005를 삽입하려면, 정렬 순서를 유지하기 위하여 1000 다음의 모든 요소들을 이동시켜야 한다.  
배열의 경우 특별한 기술을 사용하지 않는 한 삭제 또한 비용이 크다. 예를 들어, id[]에서 1010을 삭제할 경우, 1010 이후의 모든 값을 이동시켜야 한다.  

## 배열 대비 연결 리스트의 장점
1) 동적 크기
2) 삽입/삭제 용이성

## 단점
1) 랜덤 접근이 허용되지 않는다. 즉, 첫 번째 노드부터 순차적으로 요소에 접근해야 한다. 따라서 연결 리스트의 기본 구현으로는 효율적으로 이진 탐색을 할 수 없다.
2) 포인터에 대한 메모리 공간이 리스트의 각 요소에 추가되어야 한다.
3) 캐시 친화적이지 않다. 배열의 요소는 연결 리스트의 경우에는 없는 참조의 지역성이 있다.

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
# 노드 클래스
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# 연결 리스트 클래스
class LinkedList:
    def __init__(self):
        self.head = None
        
    # 시작 부분에 새로운 노드를 삽입하는 함수
    def push(self, new_data):
        # 노드 할당 및 데이터 초기화
        new_node = Node(new_data)

        # 새 노드의 next가 현재의 head를 가리키도록 설정
        new_node.next = self.head

        # head가 새 노드를 가리키도록 설정
        self.head = new_node
        
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

    # 연결 리스트의 내용을 Head부터 시작해서 출력하는 함수
    def printList(self):
        temp = self.head
        while temp:
            print(temp.data)
            temp = temp.next
            
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

## Reference
- https://www.geeksforgeeks.org/data-structures/linked-list/#singlyLinkedList
- https://www.geeksforgeeks.org/linked-list-set-1-introduction/
- https://www.geeksforgeeks.org/linked-list-set-2-inserting-a-node/