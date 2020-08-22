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
