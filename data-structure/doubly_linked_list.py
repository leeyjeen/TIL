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