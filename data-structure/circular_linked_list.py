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
    
# Output:
# 11
# 2
# 56
# 12
    