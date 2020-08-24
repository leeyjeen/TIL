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

    # 연결 리스트 클래스 내부의 함수
    # 연결 리스트를 순회하며 모든 노드를 삭제하는 함수
    def deleteList(self):
        current = self.head
        while current:
            next_node = current.next
            del current.data
            current = next

    # 연결 리스트 클래스 내부의 함수
    # 연결 리스트의 전체 노드 개수를 세는 함수
    def getCount(self):
        temp = self.head
        count = 0

        while temp:
            count += 1
            temp = temp.next
        return count

    # 연결 리스트에 특정 값이 존재하는지 검사하는 함수 (반복문 이용)
    def searchByIteration(self, x):
        current = self.head

        while current != None:
            if current.data == x:
                return True
            current = current.Next
        
        return False

    # 연결 리스트에 특정 값이 존재하는지 검사하는 함수 (재귀 이용)
    def searchByRecursion(self, current, x):
        if not current:
            return False
        
        if current.data == x:
            return True

        return self.searchByRecursion(current.next, x)

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
    linked_list.deleteNode(6) # 1->7->4->None
    linked_list.deleteNodeAt(1) # 1->4->None
    linked_list.insertAfter(linked_list.head.next, 8) # 1->4->8->None
    print("Element at index 2 is : ", linked_list.getNth(2))
    linked_list.printList()
